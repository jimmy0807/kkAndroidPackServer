package bean

import "strconv"

func FetchUnPackTask() []interface{} {
	sqlString := "select channel_id ChannelID, channel_name ChannelName, status Status, apk_name ApkName, DATE_FORMAT(write_time, '%Y-%m-%d %H:%i:%s') WriteTime, DATE_FORMAT(start_time, '%Y-%m-%d %H:%i:%s') StartTime, finish_time FinishTime from channel c inner join packInfo on status = 'waiting' ORDER BY c.channel_id asc limit 0,1"

	app := new(PackageApp)
	apps := app.Type(app).QueryOne(sqlString)

	return apps
}

func FetchTimeOutBuildingPackTask() []interface{} {
	sqlString := "select channel_id ChannelID, channel_name ChannelName, status Status from channel where Status = 'building' and TIMESTAMPDIFF(Second, start_time, NOW()) > 150"

	app := new(PackageApp)
	apps := app.Type(app).QueryList(sqlString)

	return apps
}

func FetchPackageTaskbyID(channelID int64) []interface{} {
	sqlString := "select channel_id ChannelID, channel_name ChannelName, status Status, write_time WriteTime, host_name HostName, DATE_FORMAT(start_time, '%Y-%m-%d %H:%i:%s') StartTime, DATE_FORMAT(finish_time, '%Y-%m-%d %H:%i:%s') FinishTime from channel c where channel_id = " + strconv.FormatInt(channelID, 10)

	app := new(PackageApp)
	apps := app.Type(app).QueryOne(sqlString)

	return apps
}

func FetchPackDir() []interface{} {
	sqlString := "select dir Dir from packInfo ORDER BY id asc limit 0,1"

	dir := new(PackageDir)
	dirs := dir.Type(dir).QueryOne(sqlString)

	return dirs
}
