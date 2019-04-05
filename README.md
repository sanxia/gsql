# gsql
Automatically generate a specified number of databases and tables according to SQL file content

---Example---

import (

    "github.com/sanxia/gsql"

)

sqlBuilder := gsql.NewSqlBuilder()

sqlBuilder.Configs = gsql.DatabaseConfigList{

    gsql.DatabaseConfig{

        Source: gsql.DatabaseSource{

            Username: "sanxia",
            Password: "sanxia",
            Host:     "127.0.0.1:3306",

        },

        Name:  "sanxia_user",

        Count: 2,

        TableConfigs: gsql.TableConfigList{

            gsql.TableConfig{

                Name:  "sx_user_account",
                Count: 2,
                IsNew: true,
            },

            gsql.TableConfig{

                Name:  "sx_user_profile",
                Count: 2,
                IsNew: true,

            },

        },

        IsNew: true,

    },

}

sqlBuilder.Build()


