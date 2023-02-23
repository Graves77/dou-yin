package logic

import (
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/types/mysqlmanageserver"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoInfo struct {
	VideoID       int64  `gorm:"cloumn:video_id;primaryKey"`
	VideoTitle    string `gorm:"cloumn:video_title"`
	AuthorID      int64  `gorm:"cloumn:author_id"`
	CoverUrl      string `gorm:"cloumn:cover_url"`
	PlayUrl       string `gorm:"cloumn:play_url"`
	FavoriteCount int64  `gorm:"cloumn:favorite_count"`
	CommentCount  int64  `gorm:"cloumn:comment_count"`
	IsFavotite    bool
}

type GetVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListLogic {
	return &GetVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发布列表视频信息
func (l *GetVideoListLogic) GetVideoList(in *mysqlmanageserver.GetVideoListRequest) (*mysqlmanageserver.GetVideoListResponse, error) {
	var ls VideoNumLogic
	ins := mysqlmanageserver.VideoNumRequest{AuthorId: in.AuthorId}
	out, err := ls.VideoNum(&ins)
	if err != nil {
		logx.Errorf("[pkg]logic [func]GetVideoList [msg]rpc VideoNum [err]%v", err)
		return &mysqlmanageserver.GetVideoListResponse{}, nil
	}
	if out.Num == 0 {
		logx.Infof("[pkg]logic [func]GetVideoList [msg]VideoList is empty")
		return &mysqlmanageserver.GetVideoListResponse{VideoInfo: []*mysqlmanageserver.VideoInfo{}}, nil
	}

	var list []VideoInfo
	err = svc.DB.Table("video_info").Where("author_id = ?", in.AuthorId).Find(&list).Error
	if err != nil {
		logx.Errorf("[pkg]logic [func]GetVideoList [msg]gorm Find [err]%v", err)
	}

	var getVideoListResponse mysqlmanageserver.GetVideoListResponse
	var lss IsFavotiteLogic
	for _, vi := range list {

		inss := mysqlmanageserver.IsFavotiteRequest{UserId: in.UserId, VideoId: vi.VideoID}
		outss, err := lss.IsFavotite(&inss)
		if err != nil {
			logx.Errorf("[pkg]logic [func]GetVideoList [msg]rpc IsFavotite %v", err)
			return &mysqlmanageserver.GetVideoListResponse{VideoInfo: []*mysqlmanageserver.VideoInfo{}}, nil
		}
		getVideoListResponse.VideoInfo = append(getVideoListResponse.VideoInfo, &mysqlmanageserver.VideoInfo{
			VideoId:       vi.VideoID,
			VideoTitle:    vi.VideoTitle,
			AuthorId:      vi.AuthorID,
			CoverUrl:      vi.CoverUrl,
			PlayUrl:       vi.PlayUrl,
			FavoriteCount: vi.FavoriteCount,
			CommentCount:  vi.CommentCount,
			IsFavotite:    outss.Ok,
		})
	}
	return &getVideoListResponse, nil

}
