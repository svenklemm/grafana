package metrics

import (
	"time"

	"github.com/grafana/grafana/pkg/bus"
	"github.com/grafana/grafana/pkg/models"
)

var (
	M_Instance_Start                       Counter
	M_Page_Status_200                      Counter
	M_Page_Status_500                      Counter
	M_Page_Status_404                      Counter
	M_Page_Status_Unknown                  Counter
	M_Api_Status_200                       Counter
	M_Api_Status_404                       Counter
	M_Api_Status_500                       Counter
	M_Api_Status_Unknown                   Counter
	M_Proxy_Status_200                     Counter
	M_Proxy_Status_404                     Counter
	M_Proxy_Status_500                     Counter
	M_Proxy_Status_Unknown                 Counter
	M_Api_User_SignUpStarted               Counter
	M_Api_User_SignUpCompleted             Counter
	M_Api_User_SignUpInvite                Counter
	M_Api_Dashboard_Save                   Timer
	M_Api_Dashboard_Get                    Timer
	M_Api_Dashboard_Search                 Timer
	M_Api_Admin_User_Create                Counter
	M_Api_Login_Post                       Counter
	M_Api_Login_OAuth                      Counter
	M_Api_Org_Create                       Counter
	M_Api_Dashboard_Snapshot_Create        Counter
	M_Api_Dashboard_Snapshot_External      Counter
	M_Api_Dashboard_Snapshot_Get           Counter
	M_Models_Dashboard_Insert              Counter
	M_Alerting_Result_State_Alerting       Counter
	M_Alerting_Result_State_Ok             Counter
	M_Alerting_Result_State_Paused         Counter
	M_Alerting_Result_State_NoData         Counter
	M_Alerting_Result_State_Pending        Counter
	M_Alerting_Notification_Sent_Slack     Counter
	M_Alerting_Notification_Sent_Email     Counter
	M_Alerting_Notification_Sent_Webhook   Counter
	M_Alerting_Notification_Sent_DingDing  Counter
	M_Alerting_Notification_Sent_PagerDuty Counter
	M_Alerting_Notification_Sent_LINE      Counter
	M_Alerting_Notification_Sent_Victorops Counter
	M_Alerting_Notification_Sent_OpsGenie  Counter
	M_Alerting_Notification_Sent_Telegram  Counter
	M_Alerting_Notification_Sent_Threema   Counter
	M_Alerting_Notification_Sent_Sensu     Counter
	M_Alerting_Notification_Sent_Pushover  Counter
	M_Aws_CloudWatch_GetMetricStatistics   Counter
	M_Aws_CloudWatch_ListMetrics           Counter
	M_DB_DataSource_QueryById              Counter

	// Timers
	M_DataSource_ProxyReq_Timer Timer
	M_Alerting_Execution_Time   Timer

	// StatTotals
	M_Alerting_Active_Alerts Gauge
	M_StatTotal_Dashboards   Gauge
	M_StatTotal_Users        Gauge
	M_StatTotal_Orgs         Gauge
	M_StatTotal_Playlists    Gauge
)

var UseNilMetrics bool

func Init(settings *MetricSettings, clients MetricFactories) {
	UseNilMetrics = settings.Enabled == false

	M_Instance_Start = clients.RegCounter("instance_start")

	M_Page_Status_200 = clients.RegCounter("page.resp_status", "code", "200")
	M_Page_Status_500 = clients.RegCounter("page.resp_status", "code", "500")
	M_Page_Status_404 = clients.RegCounter("page.resp_status", "code", "404")
	M_Page_Status_Unknown = clients.RegCounter("page.resp_status", "code", "unknown")

	M_Api_Status_200 = clients.RegCounter("api.resp_status", "code", "200")
	M_Api_Status_404 = clients.RegCounter("api.resp_status", "code", "404")
	M_Api_Status_500 = clients.RegCounter("api.resp_status", "code", "500")
	M_Api_Status_Unknown = clients.RegCounter("api.resp_status", "code", "unknown")

	M_Proxy_Status_200 = clients.RegCounter("proxy.resp_status", "code", "200")
	M_Proxy_Status_404 = clients.RegCounter("proxy.resp_status", "code", "404")
	M_Proxy_Status_500 = clients.RegCounter("proxy.resp_status", "code", "500")
	M_Proxy_Status_Unknown = clients.RegCounter("proxy.resp_status", "code", "unknown")

	M_Api_User_SignUpStarted = clients.RegCounter("api.user.signup_started")
	M_Api_User_SignUpCompleted = clients.RegCounter("api.user.signup_completed")
	M_Api_User_SignUpInvite = clients.RegCounter("api.user.signup_invite")

	M_Api_Dashboard_Save = clients.RegTimer("api.dashboard.save")
	M_Api_Dashboard_Get = clients.RegTimer("api.dashboard.get")
	M_Api_Dashboard_Search = clients.RegTimer("api.dashboard.search")

	M_Api_Admin_User_Create = clients.RegCounter("api.admin.user_create")
	M_Api_Login_Post = clients.RegCounter("api.login.post")
	M_Api_Login_OAuth = clients.RegCounter("api.login.oauth")
	M_Api_Org_Create = clients.RegCounter("api.org.create")

	M_Api_Dashboard_Snapshot_Create = clients.RegCounter("api.dashboard_snapshot.create")
	M_Api_Dashboard_Snapshot_External = clients.RegCounter("api.dashboard_snapshot.external")
	M_Api_Dashboard_Snapshot_Get = clients.RegCounter("api.dashboard_snapshot.get")

	M_Models_Dashboard_Insert = clients.RegCounter("models.dashboard.insert")

	M_Alerting_Result_State_Alerting = clients.RegCounter("alerting.result", "state", "alerting")
	M_Alerting_Result_State_Ok = clients.RegCounter("alerting.result", "state", "ok")
	M_Alerting_Result_State_Paused = clients.RegCounter("alerting.result", "state", "paused")
	M_Alerting_Result_State_NoData = clients.RegCounter("alerting.result", "state", "no_data")
	M_Alerting_Result_State_Pending = clients.RegCounter("alerting.result", "state", "pending")

	M_Alerting_Notification_Sent_Slack = clients.RegCounter("alerting.notifications_sent", "type", "slack")
	M_Alerting_Notification_Sent_Email = clients.RegCounter("alerting.notifications_sent", "type", "email")
	M_Alerting_Notification_Sent_Webhook = clients.RegCounter("alerting.notifications_sent", "type", "webhook")
	M_Alerting_Notification_Sent_DingDing = clients.RegCounter("alerting.notifications_sent", "type", "dingding")
	M_Alerting_Notification_Sent_PagerDuty = clients.RegCounter("alerting.notifications_sent", "type", "pagerduty")
	M_Alerting_Notification_Sent_Victorops = clients.RegCounter("alerting.notifications_sent", "type", "victorops")
	M_Alerting_Notification_Sent_OpsGenie = clients.RegCounter("alerting.notifications_sent", "type", "opsgenie")
	M_Alerting_Notification_Sent_Telegram = clients.RegCounter("alerting.notifications_sent", "type", "telegram")
	M_Alerting_Notification_Sent_Threema = clients.RegCounter("alerting.notifications_sent", "type", "threema")
	M_Alerting_Notification_Sent_Sensu = clients.RegCounter("alerting.notifications_sent", "type", "sensu")
	M_Alerting_Notification_Sent_LINE = clients.RegCounter("alerting.notifications_sent", "type", "LINE")
	M_Alerting_Notification_Sent_Pushover = clients.RegCounter("alerting.notifications_sent", "type", "pushover")

	M_Aws_CloudWatch_GetMetricStatistics = clients.RegCounter("aws.cloudwatch.get_metric_statistics")
	M_Aws_CloudWatch_ListMetrics = clients.RegCounter("aws.cloudwatch.list_metrics")

	M_DB_DataSource_QueryById = RegCounter("db.datasource.query_by_id")

	// Timers
	M_DataSource_ProxyReq_Timer = clients.RegTimer("api.dataproxy.request.all")
	M_Alerting_Execution_Time = clients.RegTimer("alerting.execution_time")

	// StatTotals
	M_Alerting_Active_Alerts = clients.RegGauge("alerting.active_alerts")
	M_StatTotal_Dashboards = clients.RegGauge("stat_totals", "stat", "dashboards")
	M_StatTotal_Users = clients.RegGauge("stat_totals", "stat", "users")
	M_StatTotal_Orgs = clients.RegGauge("stat_totals", "stat", "orgs")
	M_StatTotal_Playlists = clients.RegGauge("stat_totals", "stat", "playlists")

	go initStats(settings)
}

func initStats(settings *MetricSettings) {
	if !settings.Enabled {
		return
	}

	onceEveryDayTick := time.NewTicker(time.Hour * 24)
	secondTicker := time.NewTicker(time.Second * time.Duration(settings.IntervalSeconds))

	for {
		select {
		case <-onceEveryDayTick.C:
			sendUsageStats()
		case <-secondTicker.C:
			updateTotalStats()
		}
	}
}

var metricPublishCounter int64 = 0

func updateTotalStats() {
	metricPublishCounter++
	if metricPublishCounter%10 == 0 {
		statsQuery := models.GetSystemStatsQuery{}
		if err := bus.Dispatch(&statsQuery); err != nil {
			metricsLogger.Error("Failed to get system stats", "error", err)
			return
		}

		M_StatTotal_Dashboards.Update(statsQuery.Result.DashboardCount)
		M_StatTotal_Users.Update(statsQuery.Result.UserCount)
		M_StatTotal_Playlists.Update(statsQuery.Result.PlaylistCount)
		M_StatTotal_Orgs.Update(statsQuery.Result.OrgCount)
	}
}
