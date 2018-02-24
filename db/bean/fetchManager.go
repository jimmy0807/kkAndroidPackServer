package bean

func FetchUnPackTask() []interface{} {
	sqlString := "select channel_id ChannelID, channel_name ChannelName, status Status, apk_name ApkName from channel c left join packInfo on status = 'waiting' ORDER BY c.channel_id asc limit 0,1"

	app := new(PackageApp)
	apps := app.Type(app).QueryOne(sqlString)

	return apps
}
