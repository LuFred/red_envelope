create database red_envelope;
USE red_envelope;

#银行卡 [模拟用户银行余额]
drop table if EXISTS bank_card;
create TABLE if not EXISTS bank_card(
`id` int  auto_increment,#自增id
`user_id` int not null,#用户id
`money` int not null default 0,#银行卡金额 精确到分的整数
`gmt_create` bigint NOT NULL,  
`gmt_modified` bigint not null default 0,
PRIMARY KEY (`id`)
)ENGINE=INNODB DEFAULT charset=utf8;

#红包表 
drop table if EXISTS red_envelope;
create TABLE if not EXISTS red_envelope(
`id` int  auto_increment,#自增id
`user_id` int not null,#红包创建用户id
`secret_code` varchar(100) not null ,#红包口令
`amount` int not null ,#红包金额 精确到分的整数
`count` int not null,#红包个数
`expire_time` bigint not null,#到期时间 毫秒为单位 
`gmt_create` bigint NOT NULL,  
`gmt_modified` bigint not null default 0,
PRIMARY KEY (`id`)
)ENGINE=INNODB DEFAULT charset=utf8;
#红包领取记录
drop table if EXISTS receive_record;
create TABLE if not EXISTS receive_record(
`id` int  auto_increment,#自增id
`user_id` int not null,#领取人用户id
`red_envelope_id` int not null ,#红包id
`amount` int not null ,#领取金额 精确到分的整数
`gmt_create` bigint NOT NULL,  
`gmt_modified` bigint not null default 0,
PRIMARY KEY (`id`)
)ENGINE=INNODB DEFAULT charset=utf8;

#红包过期清退表 
drop table if EXISTS red_envelope;
create TABLE if not EXISTS red_envelope(
`id` int  auto_increment,#自增id
`user_id` int not null,#红包创建用户id
`red_envelope_id` int not null ,#红包id
`amount` int not null ,#待清退金额 精确到分的整数
`status` TINYINT(1) not null,#处理状态 0：未处理|1：已成功清退|2：清退失败
`gmt_create` bigint NOT NULL,  
`gmt_modified` bigint not null default 0,
PRIMARY KEY (`id`)
)ENGINE=INNODB DEFAULT charset=utf8;


#余额
drop table if EXISTS balance;
create TABLE if not EXISTS balance(
`id` int  auto_increment,#自增id
`user_id` int not null,#用户id
`balance` int not null ,#红包金额 精确到分的整数
`gmt_create` bigint NOT NULL,  
`gmt_modified` bigint not null default 0,
PRIMARY KEY (`id`)
)ENGINE=INNODB DEFAULT charset=utf8;
