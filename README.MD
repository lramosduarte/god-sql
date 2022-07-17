`god-sql` is a really simple sql builder and nothing more!

> Disclaimer: this repository was created sometime in the past from scratch for me to learn a little about golang. Now I decided to push and maybe improve and/or finish some things.

# Features
- [X] SQL Builder
- [X] Separeted builder `Select`(DQL) statement available
    - predicate(`where` clause) builder
    - Select columns builder only
- [X] compatibilty with sql `ansi-92`
- [] DML Builder
- [] Building query(returned columns or predicates) from `maps`, `struct` or `slice`

# What ins't expected to be done
- Parameter interpolation

# Builder

Select:
```golang
import github.com/lramosduarte/god-sql/builder

...
    sql := builder.Builder{}.Select().From("foo_bar").Build()
...
```
output:
> SELECT * FROM foo_bar

Select With filters:
```golang
import github.com/lramosduarte/god-sql/builder

...
    ps := predicate.Predicates{}
    ps.And("c1", predicate.Equal, "$1").And("c2", predicate.Greater, "$2")
    sql := builder.Builder{}.Select().From("foo_bar").Where(ps).Build()
...
```
output:
> SELECT * FROM foo_bar WHERE c1 = $1 AND c2 > $2

Select specified columns:
```golang
import github.com/lramosduarte/god-sql/builder

...
    sql := builder.Builder{}.Select("foo", "bar").From("foo_bar").Where().Build()
...
```
output:
> SELECT foo, bar FROM foo_bar WHERE c1 = $1 AND c2 > $2




**More samples in test files.**