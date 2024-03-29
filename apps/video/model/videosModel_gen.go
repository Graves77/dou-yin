// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	videosFieldNames          = builder.RawFieldNames(&Videos{})
	videosRows                = strings.Join(videosFieldNames, ",")
	videosRowsExpectAutoSet   = strings.Join(stringx.Remove(videosFieldNames, "`video_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	videosRowsWithPlaceHolder = strings.Join(stringx.Remove(videosFieldNames, "`video_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheDouyinVideoVideosVideoIdPrefix = "cache:douyinVideo:videos:videoId:"
)

type (
	videosModel interface {
		Insert(ctx context.Context, data *Videos) (sql.Result, error)
		FindOne(ctx context.Context, videoId int64) (*Videos, error)
		FindListByUid(ctx context.Context, userId int64) ([]*Videos, error)
		Update(ctx context.Context, data *Videos) error
		Delete(ctx context.Context, videoId int64) error
	}

	defaultVideosModel struct {
		sqlc.CachedConn
		table string
	}

	Videos struct {
		VideoId       int64  `db:"video_id"`       // 视频唯一标识
		UserId        int64  `db:"user_id"`        // 视频作者id
		PlayUrl       string `db:"play_url"`       // 视频播放地址
		CoverUrl      string `db:"cover_url"`      // 视频封面地址
		FavoriteCount int64  `db:"favorite_count"` // 视频的点赞总数
		CommentCount  int64  `db:"comment_count"`  // 视频的评论总数
		Title         string `db:"title"`          // 视频标题
	}
)

func newVideosModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultVideosModel {
	return &defaultVideosModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`videos`",
	}
}

func (m *defaultVideosModel) Delete(ctx context.Context, videoId int64) error {
	douyinVideoVideosVideoIdKey := fmt.Sprintf("%s%v", cacheDouyinVideoVideosVideoIdPrefix, videoId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `video_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, videoId)
	}, douyinVideoVideosVideoIdKey)
	return err
}

func (m *defaultVideosModel) FindOne(ctx context.Context, videoId int64) (*Videos, error) {
	douyinVideoVideosVideoIdKey := fmt.Sprintf("%s%v", cacheDouyinVideoVideosVideoIdPrefix, videoId)
	var resp Videos
	err := m.QueryRowCtx(ctx, &resp, douyinVideoVideosVideoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `video_id` = ? limit 1", videosRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, videoId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVideosModel) FindListByUid(ctx context.Context, userId int64) ([]*Videos, error) {
	var resp []*Videos
	err := m.QueryRowsNoCacheCtx(ctx, &resp, fmt.Sprintf("select %s from %s where `author_id` = ?", videosRows, m.table), userId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultVideosModel) Insert(ctx context.Context, data *Videos) (sql.Result, error) {
	douyinVideoVideosVideoIdKey := fmt.Sprintf("%s%v", cacheDouyinVideoVideosVideoIdPrefix, data.VideoId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, videosRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.PlayUrl, data.CoverUrl, data.FavoriteCount, data.CommentCount, data.Title)
	}, douyinVideoVideosVideoIdKey)
	return ret, err
}

func (m *defaultVideosModel) Update(ctx context.Context, data *Videos) error {
	douyinVideoVideosVideoIdKey := fmt.Sprintf("%s%v", cacheDouyinVideoVideosVideoIdPrefix, data.VideoId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `video_id` = ?", m.table, videosRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.PlayUrl, data.CoverUrl, data.FavoriteCount, data.CommentCount, data.Title, data.VideoId)
	}, douyinVideoVideosVideoIdKey)
	return err
}

func (m *defaultVideosModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheDouyinVideoVideosVideoIdPrefix, primary)
}

func (m *defaultVideosModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `video_id` = ? limit 1", videosRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultVideosModel) tableName() string {
	return m.table
}
