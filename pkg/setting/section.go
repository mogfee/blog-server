package setting

import "time"

type ServiceSettingS struct {
	RunModel     string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
type AppSettingS struct {
	ServerName      string
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}
type DatabaseSettingS struct {
	DBType          string
	Username        string
	Password        string
	Host            string
	DBName          string
	TablePrefix     string
	Charset         string
	ParseTime       bool
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type JaegerSettingS struct {
	Host string
	Port int64
}

type MonitorSettingS struct {
	SystemEmailUser string
	SystemEmailPass string
	SystemEmailHost string
	SystemEmailPort int
	ErrorNotifyUser string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	if err := s.vp.UnmarshalKey(k, v); err != nil {
		return err
	}
	return nil
}
