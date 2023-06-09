package container

import (
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/ropel12/email/config"
	"github.com/ropel12/email/pkg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Depend struct {
	Config      *config.Config
	NSQConsumer *pkg.NSQConsumer
	Db          *gorm.DB
}

func InitContainer() (*Depend, error) {
	config, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	nsqConsumer, err := NewNSQConsumer(config)
	if err != nil {
		return nil, err
	}
	db, err := GetConnection(config)
	if err != nil {
		return nil, err
	}
	return &Depend{
		Config:      config,
		NSQConsumer: nsqConsumer,
		Db:          db,
	}, nil
}
func NewNSQConsumer(conf *config.Config) (*pkg.NSQConsumer, error) {
	nc := &pkg.NSQConsumer{}
	nc.Env = conf.NSQ
	var err error
	nsqConfig := nsq.NewConfig()
	nc.Consumer, err = nsq.NewConsumer(nc.Env.Topic, nc.Env.Channel, nsqConfig)
	if err != nil {
		return nil, err
	}

	nc.Consumer2, err = nsq.NewConsumer(nc.Env.Topic2, nc.Env.Channel2, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer3, err = nsq.NewConsumer(nc.Env.Topic3, nc.Env.Channel3, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer4, err = nsq.NewConsumer(nc.Env.Topic4, nc.Env.Channel4, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer5, err = nsq.NewConsumer(nc.Env.Topic5, nc.Env.Channel5, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer6, err = nsq.NewConsumer(nc.Env.Topic6, nc.Env.Channel6, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer7, err = nsq.NewConsumer(nc.Env.Topic7, nc.Env.Channel7, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer8, err = nsq.NewConsumer(nc.Env.Topic8, nc.Env.Channel8, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer9, err = nsq.NewConsumer(nc.Env.Topic9, nc.Env.Channel8, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer10, err = nsq.NewConsumer(nc.Env.Topic10, nc.Env.Channel10, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer11, err = nsq.NewConsumer(nc.Env.Topic11, nc.Env.Channel10, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer12, err = nsq.NewConsumer(nc.Env.Topic12, nc.Env.Channel10, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer13, err = nsq.NewConsumer(nc.Env.Topic13, nc.Env.Channel10, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer14, err = nsq.NewConsumer(nc.Env.Topic14, nc.Env.Channel10, nsqConfig)
	if err != nil {
		return nil, err
	}
	nc.Consumer15, err = nsq.NewConsumer(nc.Env.Topic15, nc.Env.Channel10, nsqConfig)
	if err != nil {
		return nil, err
	}
	return nc, nil
}
func GetConnection(c *config.Config) (*gorm.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	)
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("unable to access database sql: %v", err)
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("unable to establish a good connection to the database: %v", err)
	}

	return db, nil
}
