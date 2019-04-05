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


----- Builder Result -----
=====
Create Database:
=====
    1. sanxia_user_00000
    -----
    Create Table:
    -----
    1.1 sx_user_account_00000
    1.2 sx_user_account_00001
    1.3 sx_user_profile_00000
    1.4 sx_user_profile_00001

    2. sanxia_user_00001
    -----
    Create Table:
    -----
    2.1 sx_user_account_00000
    2.2 sx_user_account_00001
    2.3 sx_user_profile_00000
    2.4 sx_user_profile_00001


