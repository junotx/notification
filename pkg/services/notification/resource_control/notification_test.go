// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.
package resource_control

import (
	"testing"

	"openpitrix.io/logger"

	pkg "openpitrix.io/notification/pkg"
	"openpitrix.io/notification/pkg/config"
	"openpitrix.io/notification/pkg/models"
	"openpitrix.io/notification/pkg/pb"
	"openpitrix.io/notification/pkg/util/pbutil"
)

func TestDescribeNotifications4rc(t *testing.T) {
	if !*pkg.LocalDevEnvEnabled {
		t.Skip("Local Dev testing env disabled.")
	}

	config.GetInstance()

	var nfIds []string
	nfIds = append(nfIds, "nf-yM793AqkEmnj")
	nfIds = append(nfIds, "nf-lLZ9L8OzZwnj")

	var contentTypes []string
	contentTypes = append(contentTypes, "email")

	var owners []string
	owners = append(owners, "HuoJiao")

	var statuses []string
	statuses = append(statuses, "successful")

	var req = &pb.DescribeNotificationsRequest{
		NotificationId: nfIds,
		ContentType:    contentTypes,
		Owner:          owners,
		Status:         statuses,
		Limit:          20,
		Offset:         0,
		SearchWord:     pbutil.ToProtoString("successful"),
		SortKey:        pbutil.ToProtoString("status"),
		Reverse:        pbutil.ToProtoBool(false),
		DisplayColumns: nil,
	}

	notifications, cnt, err := DescribeNotifications(nil, req)

	if err != nil {
		logger.Errorf(nil, "Failed to describe notifications, error, %+v.", err)
	}

	logger.Infof(nil, "Test describe notifications:,cnt = %d，notifications=[%+v]", cnt, notifications)
}

func TestRegisterNotification4rc(t *testing.T) {
	if !*pkg.LocalDevEnvEnabled {
		t.Skip("Local Dev testing env disabled.")
	}
	config.GetInstance()
	testAddrsStr := "{\"email\": [\"openpitrix@163.com\", \"openpitrix@163.com\"]}"
	//testAddrListIds := "[\"adl-EgoLADQkwkEr\"]"
	testExtra := "{\"ws_service\": \"ks\",\"ws_message_type\": \"event\"}"
	var req = &pb.CreateNotificationRequest{
		ContentType:        pbutil.ToProtoString("alert"),
		Title:              pbutil.ToProtoString("testing alert"),
		Content:            pbutil.ToProtoString("test content"),
		ShortContent:       pbutil.ToProtoString("test short content"),
		ExpiredDays:        pbutil.ToProtoUInt32(0),
		Owner:              pbutil.ToProtoString("HuoJiao"),
		AddressInfo:        pbutil.ToProtoString(testAddrsStr),
		AvailableStartTime: pbutil.ToProtoString(""),
		AvailableEndTime:   pbutil.ToProtoString(""),
		Extra:              pbutil.ToProtoString(testExtra),
	}

	notification := models.NewNotification(req)

	err := RegisterNotification(nil, notification)
	if err != nil {
		logger.Errorf(nil, "Failed to register notification, %+v.", err)
	}
	logger.Debugf(nil, "RegisterNotification [%s] in DB successfully.", notification.NotificationId)

}
