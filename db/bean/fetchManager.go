package bean

import "strconv"

func FetchUnPackTask() []interface{} {
	sqlString := "select channel_id ChannelID, channel_name ChannelName, status Status, apk_name ApkName, write_time WriteTime from channel c inner join packInfo on status = 'waiting' ORDER BY c.channel_id asc limit 0,1"

	app := new(PackageApp)
	apps := app.Type(app).QueryOne(sqlString)

	return apps
}

func FetchPackageTaskbyID(channelID int64) []interface{} {
	sqlString := "select channel_id ChannelID, channel_name ChannelName, status Status, write_time WriteTime, host_name HostName from channel c where channel_id = " + strconv.FormatInt(channelID, 10)

	app := new(PackageApp)
	apps := app.Type(app).QueryOne(sqlString)

	return apps
}
