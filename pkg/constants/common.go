// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package constants

const (
	EtcdPrefix      = "notification/"
	EmailQueue      = "emailtask"
	MaxWorkingTasks = 5
)

const (
	NfPostIDPrifix   = "nf-"
	JobPostIDPrifix  = "job-"
	TaskPostIDPrifix = "task-"
)

const (
	StatusNew      = "new"
	StatusSending  = "sending"
	StatusFinished = "finished"
	StatusFailed   = "failed"
)

const (
	ServiceName    = "Notification"
	prefix         = "notification-"
	ApiGatewayHost = prefix + "api-gateway"
	//ApiGatewayHost = "127.0.0.1"
	ApiGatewayPort = 9200

	NotificationManagerHost = prefix + "manager"
	//NotificationManagerHost = "127.0.0.1"
	//NotificationManagerHost = "192.168.0.3"
	NotificationManagerPort = 9201
)