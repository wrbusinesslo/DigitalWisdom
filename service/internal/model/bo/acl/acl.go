package boAcl

import aclDaoModel "DigitalWisdom/service/dao/daoModels/acl"

type CreateAccountArgs struct {
	Query *aclDaoModel.Query
}

type AclArgs struct {
	Query *aclDaoModel.AccountPassword
}
