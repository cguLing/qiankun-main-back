package utils

import (
	"bus-backend-go/conf"
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"strings"
	"time"
)

var Log = logrus.New()
func init() {
}
//cfg 配置文件
type cfg struct {
	LogLvl     string   // 日志级别
	EsAddrs    []string //ES addr
	EsUser     string   //ES user
	EsPassword string   //ES password
}

//setupLogrus 初始化logrus 同时把logrus的logger var 引用到这个common.Logger
func setupLogrus(cc cfg, log *logrus.Logger) error {
	logLvl, err := logrus.ParseLevel(cc.LogLvl)
	if err != nil {
		return err
	}
	log.SetLevel(logLvl)
	//log.Println(log.GetLevel())
	log.SetReportCaller(true)
	//开启 logrus ES hook
	esh := newEsHook(cc)
	log.AddHook(esh)
	return nil
}
func New(log *logrus.Logger) {
	cc := cfg{
		LogLvl:     "info",
		EsAddrs:    conf.Sysconfig.LogAddr, //[]string{"http://localhost:9200/"},
		EsUser:     conf.Sysconfig.LogUsername,
		EsPassword: conf.Sysconfig.LogPass,
	}
	err := setupLogrus(cc, log)
	if err != nil {
		log.Println(err)
	}
}

//esHook 自定义的ES hook
type esHook struct {
	cmd    string // 记录启动的命令
	client *elastic.Client
}

//newEsHook 初始化
func newEsHook(cc cfg) *esHook {
	if cc.EsUser == "" && cc.EsPassword == "" {
		es, err := elastic.NewClient(
			elastic.SetURL(cc.EsAddrs...),
			//elastic.SetBasicAuth(cc.EsUser, cc.EsPassword),
			elastic.SetSniff(false),
			elastic.SetHealthcheckInterval(15*time.Second),
			elastic.SetErrorLog(log.New(os.Stderr, "ES:", log.LstdFlags)),
		)
		if err != nil {
			log.Println("failed to create Elastic V6 Client: ", err)
		}
		return &esHook{client: es, cmd: strings.Join(os.Args, " ")}
	} else {
		es, err := elastic.NewClient(
			elastic.SetURL(cc.EsAddrs...),
			elastic.SetBasicAuth(cc.EsUser, cc.EsPassword),
			elastic.SetSniff(false),
			elastic.SetHealthcheckInterval(15*time.Second),
			elastic.SetErrorLog(log.New(os.Stderr, "ES:", log.LstdFlags)),
		)
		if err != nil {
			log.Println("failed to create Elastic V6 Client: ", err)
		}
		return &esHook{client: es, cmd: strings.Join(os.Args, " ")}
	}

}

//Fire logrus hook interface 方法
func (hook *esHook) Fire(entry *logrus.Entry) error {
	doc := newEsLog(entry)
	doc["cmd"] = hook.cmd
	go hook.sendEs(doc)
	return nil
}

//Levels logrus hook interface 方法
func (hook *esHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.InfoLevel,
	}
}

//sendEs 异步发送日志到es
func (hook *esHook) sendEs(doc appLogDocModel) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("send entry to es failed: ", r)
		}
	}()
	_, err := hook.client.Index().Index(doc.indexName()).Type("_doc").BodyJson(doc).Do(context.Background())
	if err != nil {
		log.Println(err)
	}

}

//appLogDocModel es model
type appLogDocModel map[string]interface{}

func newEsLog(e *logrus.Entry) appLogDocModel {
	ins := map[string]interface{}{}
	for kk, vv := range e.Data {
		ins[kk] = vv
	}
	ins["time"] = time.Now().Local()
	ins["level"] = e.Level
	ins["message"] = e.Message
	ins["caller"] = fmt.Sprintf("%s:%d  %#v", e.Caller.File, e.Caller.Line, e.Caller.Func)
	return ins
}

// indexName es index name 时间分割
func (m *appLogDocModel) indexName() string {
	return conf.Sysconfig.LogIndex + "." + time.Now().Local().Format("2006.01.02")
}

