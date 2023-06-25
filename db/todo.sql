CREATE TABLE bubble.todo
(
    id          bigint AUTO_INCREMENT COMMENT 'id' PRIMARY KEY,
    title       varchar(256)                       NOT NULL COMMENT '标题',
    status      boolean                            NOT NULL DEFAULT FALSE COMMENT '状态(是否完成)',
    create_time datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间',
    update_time datetime DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    is_delete   tinyint  DEFAULT 0                 NOT NULL COMMENT '是否删除'
)
    COMMENT '待办事项';
