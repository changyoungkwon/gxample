package database

import "github.com/changyoungkwon/gxample/internal/config"

var databaseURL = config.Get().Database.URL
