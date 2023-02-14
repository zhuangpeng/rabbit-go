import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	globalkey "github.com/zhuangpeng/rabbit-go/pkg/globalKey"
	"github.com/zhuangpeng/rabbit-go/pkg/statuserr"
	"github.com/zhuangpeng/rabbit-go/pkg/i18n"
	"github.com/zhuangpeng/rabbit-go/pkg/utils/dbx"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)
