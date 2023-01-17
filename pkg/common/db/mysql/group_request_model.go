package mysql

//
//func UpdateGroupRequest(groupRequest GroupRequest) error {
//	if groupRequest.HandledTime.Unix() < 0 {
//		groupRequest.HandledTime = utils.UnixSecondToTime(0)
//	}
//	return db.DB.MysqlDB.DefaultGormDB().Table("group_requests").Where("group_id=? and user_id=?", groupRequest.GroupID, groupRequest.UserID).Updates(&groupRequest).Error
//}
//
//func InsertIntoGroupRequest(toInsertInfo GroupRequest) error {
//	DelGroupRequestByGroupIDAndUserID(toInsertInfo.GroupID, toInsertInfo.UserID)
//	if toInsertInfo.HandledTime.Unix() < 0 {
//		toInsertInfo.HandledTime = utils.UnixSecondToTime(0)
//	}
//	u := db.DB.MysqlDB.DefaultGormDB().Table("group_requests").Where("group_id=? and user_id=?", toInsertInfo.GroupID, toInsertInfo.UserID).Updates(&toInsertInfo)
//	if u.RowsAffected != 0 {
//		return nil
//	}
//
//	toInsertInfo.ReqTime = time.Now()
//	if toInsertInfo.HandledTime.Unix() < 0 {
//		toInsertInfo.HandledTime = utils.UnixSecondToTime(0)
//	}
//
//	err := db.DB.MysqlDB.DefaultGormDB().Table("group_requests").Create(&toInsertInfo).Error
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func GetGroupRequestByGroupIDAndUserID(groupID, userID string) (*GroupRequest, error) {
//	var groupRequest GroupRequest
//	err := db.DB.MysqlDB.DefaultGormDB().Table("group_requests").Where("user_id=? and group_id=?", userID, groupID).Take(&groupRequest).Error
//	if err != nil {
//		return nil, err
//	}
//	return &groupRequest, nil
//}
//
//func DelGroupRequestByGroupIDAndUserID(groupID, userID string) error {
//	return db.DB.MysqlDB.DefaultGormDB().Table("group_requests").Where("group_id=? and user_id=?", groupID, userID).Delete(GroupRequest{}).Error
//}
//
//func GetGroupRequestByGroupID(groupID string) ([]GroupRequest, error) {
//	var groupRequestList []GroupRequest
//	err := db.DB.MysqlDB.DefaultGormDB().Table("group_requests").Where("group_id=?", groupID).Find(&groupRequestList).Error
//	if err != nil {
//		return nil, err
//	}
//	return groupRequestList, nil
//}
//
//received
//func GetRecvGroupApplicationList(userID string) ([]GroupRequest, error) {
//	var groupRequestList []GroupRequest
//	memberList, err := GetGroupMemberListByUserID(userID)
//	if err != nil {
//		return nil, utils.Wrap(err, utils.GetSelfFuncName())
//	}
//	for _, v := range memberList {
//		if v.RoleLevel > constant.GroupOrdinaryUsers {
//			list, err := GetGroupRequestByGroupID(v.GroupID)
//			if err != nil {
//				continue
//			}
//			groupRequestList = append(groupRequestList, list...)
//		}
//	}
//	return groupRequestList, nil
//}
//
//func GetUserReqGroupByUserID(userID string) ([]GroupRequest, error) {
//	var groupRequestList []GroupRequest
//	err := db.DB.MysqlDB.DefaultGormDB().Table("group_requests").Where("user_id=?", userID).Find(&groupRequestList).Error
//	return groupRequestList, err
//}
//
//
//func GroupApplicationResponse(pb *group.GroupApplicationResponseReq) (*group.CommonResp, error) {
//
//	ownerUser, err := FindGroupMemberInfoByGroupIdAndUserId(pb.GroupID, pb.OwnerID)
//	if err != nil {
//		log.ErrorByKv("FindGroupMemberInfoByGroupIdAndUserId failed", pb.OperationID, "groupId", pb.GroupID, "ownerID", pb.OwnerID)
//		return nil, err
//	}
//	if ownerUser.AdministratorLevel <= 0 {
//		return nil, errors.New("insufficient permissions")
//	}
//
//	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
//	if err != nil {
//		return nil, err
//	}
//	var groupRequest GroupRequest
//	err = dbConn.Raw("select * from `group_request` where handled_user = ? and group_id = ? and from_user_id = ? and to_user_id = ?",
//		"", pb.GroupID, pb.FromUserID, pb.ToUserID).Scan(&groupRequest).Error
//	if err != nil {
//		log.ErrorByKv("find group_request info failed", pb.OperationID, "groupId", pb.GroupID, "fromUserId", pb.FromUserID, "toUserId", pb.OwnerID)
//		return nil, err
//	}
//
//	if groupRequest.Flag != 0 {
//		return nil, errors.New("application has already handle")
//	}
//
//	var saveFlag int
//	if pb.HandleResult == 0 {
//		saveFlag = -1
//	} else if pb.HandleResult == 1 {
//		saveFlag = 1
//	} else {
//		return nil, errors.New("parma HandleResult error")
//	}
//	err = dbConn.Exec("update `group_request` set flag = ?, handled_msg = ?, handled_user = ? where group_id = ? and from_user_id = ? and to_user_id = ?",
//		saveFlag, pb.HandledMsg, pb.OwnerID, groupRequest.GroupID, groupRequest.FromUserID, groupRequest.ToUserID).Error
//	if err != nil {
//		log.ErrorByKv("update group request failed", pb.OperationID, "groupID", pb.GroupID, "flag", saveFlag, "ownerId", pb.OwnerID, "fromUserId", pb.FromUserID, "toUserID", pb.ToUserID)
//		return nil, err
//	}
//
//	if saveFlag == 1 {
//		if groupRequest.ToUserID == "0" {
//			err = InsertIntoGroupMember(pb.GroupID, pb.FromUserID, groupRequest.FromUserNickname, groupRequest.FromUserFaceUrl, 0)
//			if err != nil {
//				log.ErrorByKv("InsertIntoGroupMember failed", pb.OperationID, "groupID", pb.GroupID, "fromUserId", pb.FromUserID)
//				return nil, err
//			}
//		} else {
//			err = InsertIntoGroupMember(pb.GroupID, pb.ToUserID, groupRequest.ToUserNickname, groupRequest.ToUserFaceUrl, 0)
//			if err != nil {
//				log.ErrorByKv("InsertIntoGroupMember failed", pb.OperationID, "groupID", pb.GroupID, "fromUserId", pb.FromUserID)
//				return nil, err
//			}
//		}
//	}
//
//	return &group.GroupApplicationResponseResp{}, nil
//}
//
//func FindGroupBeInvitedRequestInfoByUidAndGroupID(groupId, uid string) (*GroupRequest, error) {
//	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
//	if err != nil {
//		return nil, err
//	}
//	var beInvitedRequestUserInfo GroupRequest
//	err = dbConn.Table("group_request").Where("to_user_id=? and group_id=?", uid, groupId).Find(&beInvitedRequestUserInfo).Error
//	if err != nil {
//		return nil, err
//	}
//	return &beInvitedRequestUserInfo, nil
//
//}
//
//func InsertGroupRequest(groupId, fromUser, fromUserNickName, fromUserFaceUrl, toUser, requestMsg, handledMsg string, handleStatus int) error {
//	return nil
//}