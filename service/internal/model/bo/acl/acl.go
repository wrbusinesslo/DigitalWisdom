package acl

import "DigitalWisdom/service/dao/daoModels/acl"

type GetArgs struct {
	User aclDaoModel.User
}

type GetReply struct {
	User *aclDaoModel.User
}

type GetLoginReply struct {
	Session *aclDaoModel.UserSession
}

type UpdateArgs struct {
	Query *aclDaoModel.Query
}

type UpdateReply struct{}
