# Go URL shortener service

## Schema:

Auth user -> Create short link -> Saving data in DB

## Methods:

register  |
login     |---  User

create    |
goTo      |---  Link
get       |

get       |---  Stat

## Architecture:

payload      DTO
 |
handler      http request handler
 |
repository   DAO
errors       module errors
 |
service      buisness logic

