// Code generated by config model builder generator; DO NOT EDIT.

package datasourcemodel

import (
	"encoding/json"

	tfconfig "github.com/hashicorp/terraform-plugin-testing/config"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/config"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/provider/datasources"
)

type UsersModel struct {
	Like           tfconfig.Variable `json:"like,omitempty"`
	Limit          tfconfig.Variable `json:"limit,omitempty"`
	StartsWith     tfconfig.Variable `json:"starts_with,omitempty"`
	Users          tfconfig.Variable `json:"users,omitempty"`
	WithDescribe   tfconfig.Variable `json:"with_describe,omitempty"`
	WithParameters tfconfig.Variable `json:"with_parameters,omitempty"`

	*config.DatasourceModelMeta
}

/////////////////////////////////////////////////
// Basic builders (resource name and required) //
/////////////////////////////////////////////////

func Users(
	datasourceName string,
) *UsersModel {
	u := &UsersModel{DatasourceModelMeta: config.DatasourceMeta(datasourceName, datasources.Users)}
	return u
}

func UsersWithDefaultMeta() *UsersModel {
	u := &UsersModel{DatasourceModelMeta: config.DatasourceDefaultMeta(datasources.Users)}
	return u
}

///////////////////////////////////////////////////////
// set proper json marshalling and handle depends on //
///////////////////////////////////////////////////////

func (u *UsersModel) MarshalJSON() ([]byte, error) {
	type Alias UsersModel
	return json.Marshal(&struct {
		*Alias
		DependsOn                 []string                      `json:"depends_on,omitempty"`
		SingleAttributeWorkaround config.ReplacementPlaceholder `json:"single_attribute_workaround,omitempty"`
	}{
		Alias:                     (*Alias)(u),
		DependsOn:                 u.DependsOn(),
		SingleAttributeWorkaround: config.SnowflakeProviderConfigSingleAttributeWorkaround,
	})
}

func (u *UsersModel) WithDependsOn(values ...string) *UsersModel {
	u.SetDependsOn(values...)
	return u
}

/////////////////////////////////
// below all the proper values //
/////////////////////////////////

func (u *UsersModel) WithLike(like string) *UsersModel {
	u.Like = tfconfig.StringVariable(like)
	return u
}

// limit attribute type is not yet supported, so WithLimit can't be generated

func (u *UsersModel) WithStartsWith(startsWith string) *UsersModel {
	u.StartsWith = tfconfig.StringVariable(startsWith)
	return u
}

// users attribute type is not yet supported, so WithUsers can't be generated

func (u *UsersModel) WithWithDescribe(withDescribe bool) *UsersModel {
	u.WithDescribe = tfconfig.BoolVariable(withDescribe)
	return u
}

func (u *UsersModel) WithWithParameters(withParameters bool) *UsersModel {
	u.WithParameters = tfconfig.BoolVariable(withParameters)
	return u
}

//////////////////////////////////////////
// below it's possible to set any value //
//////////////////////////////////////////

func (u *UsersModel) WithLikeValue(value tfconfig.Variable) *UsersModel {
	u.Like = value
	return u
}

func (u *UsersModel) WithLimitValue(value tfconfig.Variable) *UsersModel {
	u.Limit = value
	return u
}

func (u *UsersModel) WithStartsWithValue(value tfconfig.Variable) *UsersModel {
	u.StartsWith = value
	return u
}

func (u *UsersModel) WithUsersValue(value tfconfig.Variable) *UsersModel {
	u.Users = value
	return u
}

func (u *UsersModel) WithWithDescribeValue(value tfconfig.Variable) *UsersModel {
	u.WithDescribe = value
	return u
}

func (u *UsersModel) WithWithParametersValue(value tfconfig.Variable) *UsersModel {
	u.WithParameters = value
	return u
}
