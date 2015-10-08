# mylist

Shows Wunderlist tasks assigned to me, due today/overdue, or starred.

## Installation & Usage

First, you'll need a proper Go environment set up. Follow the instructions at [golang.org](https://golang.org/doc/install).

```
go get github.com/justincampbell/mylist
```

Grab credentials for your Wunderlist account:

1. Create a new app at [developer.wunderlist.com](https://developer.wunderlist.com/apps/new)
2. Name your app. I called mine "wl".
3. Put in a dummy URL and callback URL.
4. On the next page, click "Create Access Token". Your token will be displayed in a flash message near the top.
5. Set the environment variables `WL_ACCESS_TOKEN` and `WL_CLIENT_ID`. You can also use these with the [wl](https://github.com/robdimsdale/wl) client (which this uses). You can put these in your `~/.profile` or `~/.bashrc` so that they're available every time you open a shell.

```bash
export WL_CLIENT_ID=8783ddf34242a66f7576
export WL_ACCESS_TOKEN=59b3e5185ae95b46ec9e5a11b532e79e55e232663e6eb2bdedd3230d3e37
```

You should be able to run `mylist` to see your task list:

```
$ mylist
✅ 5
Trash cans
Straighten
Dishes
Setup Verizon
Setup electric
```
