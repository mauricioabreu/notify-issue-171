# notify-issue-171
Demo of issue 171 of https://github.com/rjeczalik/notify

# Steps

* Build the image: `docker build -t notify-issue-171 .`

* Run the program and the script that creates events: `docker run --rm -it notify-issue-171 sh /go/events.sh`

Output here:
```
2019/01/12 23:31:09.779948 [D] dispatching notify.Create on "/go/example/foo/a.txt"
2019/01/12 23:31:09.780137 New event /go/example/foo/a.txt notify.Create
2019/01/12 23:31:09.781894 [D] dispatching notify.Create on "/go/example/bar/b.txt"
2019/01/12 23:31:09.782185 New event /go/example/bar/b.txt notify.Create
2019/01/12 23:31:09.785526 [D] dispatching notify.Remove on "/go/example/foo/a.txt"
2019/01/12 23:31:09.785940 [D] dispatching notify.Remove on "/go/example/foo"
2019/01/12 23:31:09.787575 [D] dispatching notify.Remove on "/go/example/foo"
2019/01/12 23:31:09.787937 New event /go/example/foo notify.Remove
2019/01/12 23:31:09.788673 New event /go/example/foo/a.txt notify.Remove
2019/01/12 23:31:09.789476 New event /go/example/foo notify.Remove
2019/01/12 23:31:09.789367 [D] dispatching notify.Remove on "/go/example/bar/b.txt"
2019/01/12 23:31:09.790496 [D] dispatching notify.Remove on "/go/example/bar"
2019/01/12 23:31:09.790593 [D] dispatching notify.Remove on "/go/example/bar"
2019/01/12 23:31:09.790627 New event /go/example/bar notify.Remove
2019/01/12 23:31:09.790653 New event /go/example/bar/b.txt notify.Remove
2019/01/12 23:31:09.790712 New event /go/example/bar notify.Remove
2019/01/12 23:31:10.792037 [D] dispatching notify.Create on "/go/example/foo"
2019/01/12 23:31:10.792136 [D] dispatching notify.Create on "/go/example/bar"
2019/01/12 23:31:10.793051 New event /go/example/bar notify.Create
2019/01/12 23:31:10.793631 New event /go/example/foo notify.Create
2019/01/12 23:31:10.796949 [D] dispatching notify.Create on "/go/example/foo/a.txt"
2019/01/12 23:31:10.797720 New event /go/example/foo/a.txt notify.Create
2019/01/12 23:31:10.800676 [D] dispatching notify.Create on "/go/example/bar/b.txt"
2019/01/12 23:31:10.801490 New event /go/example/bar/b.txt notify.Create
2019/01/12 23:31:10.804584 [D] dispatching notify.Remove on "/go/example/foo/a.txt"
2019/01/12 23:31:10.804651 [D] dispatching notify.Remove on "/go/example/foo"
2019/01/12 23:31:10.806003 [D] dispatching notify.Remove on "/go/example/foo"
2019/01/12 23:31:10.806138 [D] dispatching notify.Remove on "/go/example/bar/b.txt"
2019/01/12 23:31:10.806154 [D] dispatching notify.Remove on "/go/example/bar"
2019/01/12 23:31:10.806692 [D] dispatching notify.Remove on "/go/example/bar"
2019/01/12 23:31:10.806934 [D] dropped notify.Remove on "/go/example/bar": receiver too slow
2019/01/12 23:31:10.806996 [D] dropped notify.Remove on "/go/example/bar/b.txt": receiver too slow
2019/01/12 23:31:10.805974 New event /go/example/foo/a.txt notify.Remove
2019/01/12 23:31:10.807123 New event /go/example/foo notify.Remove
2019/01/12 23:31:10.807132 New event /go/example/bar notify.Remove
2019/01/12 23:31:10.808071 [D] dropped notify.Remove on "/go/example/foo": receiver too slow
2019/01/12 23:31:11.813234 [D] dispatching notify.Create on "/go/example/foo"
2019/01/12 23:31:11.813316 [D] dispatching notify.Create on "/go/example/bar"
2019/01/12 23:31:11.813390 New event /go/example/bar notify.Create
2019/01/12 23:31:11.813442 New event /go/example/foo notify.Create
2019/01/12 23:31:11.819074 [D] dispatching notify.Remove on "/go/example/foo"
2019/01/12 23:31:11.819191 New event /go/example/foo notify.Remove
2019/01/12 23:31:11.819999 [D] dispatching notify.Remove on "/go/example/bar"
2019/01/12 23:31:11.820109 New event /go/example/bar notify.Remove
2019/01/12 23:31:12.826456 [D] dispatching notify.Create on "/go/example/foo"
2019/01/12 23:31:12.826534 [D] dispatching notify.Create on "/go/example/bar"
2019/01/12 23:31:12.826826 New event /go/example/bar notify.Create
2019/01/12 23:31:12.826889 New event /go/example/foo notify.Create
2019/01/12 23:31:12.836677 [D] dispatching notify.Remove on "/go/example/foo"
2019/01/12 23:31:12.836763 New event /go/example/foo notify.Remove
2019/01/12 23:31:12.837198 [D] dispatching notify.Remove on "/go/example/bar"
2019/01/12 23:31:12.837243 New event /go/example/bar notify.Remove
2019/01/12 23:31:13.846223 [D] dispatching notify.Create on "/go/example/foo"
2019/01/12 23:31:13.846386 New event /go/example/foo notify.Create
2019/01/12 23:31:13.848026 [D] dispatching notify.Create on "/go/example/bar"
2019/01/12 23:31:13.848370 New event /go/example/bar notify.Create
2019/01/12 23:31:13.854950 [D] dispatching notify.Remove on "/go/example/foo"
2019/01/12 23:31:13.855049 New event /go/example/foo notify.Remove
2019/01/12 23:31:13.856291 [D] dispatching notify.Remove on "/go/example/bar"
2019/01/12 23:31:13.857272 New event /go/example/bar notify.Remove
2019/01/12 23:31:14.861805 [D] dispatching notify.Create on "/go/example/foo"
2019/01/12 23:31:14.861898 [D] dispatching notify.Create on "/go/example/bar"
2019/01/12 23:31:14.862589 New event /go/example/bar notify.Create
2019/01/12 23:31:14.863700 New event /go/example/foo notify.Create
2019/01/12 23:31:14.869514 [D] dispatching notify.Remove on "/go/example/foo"
2019/01/12 23:31:14.869547 [D] dispatching notify.Remove on "/go/example/bar"
2019/01/12 23:31:14.869567 New event /go/example/bar notify.Remove
2019/01/12 23:31:14.869581 New event /go/example/foo notify.Remove
2019/01/12 23:31:15.875036 [D] dispatching notify.Create on "/go/example/foo"
2019/01/12 23:31:15.875510 New event /go/example/foo notify.Create
2019/01/12 23:31:15.875875 [D] dispatching notify.Create on "/go/example/bar"
2019/01/12 23:31:15.876220 New event /go/example/bar notify.Create
2019/01/12 23:31:15.885062 [D] dispatching notify.Remove on "/go/example/foo"
2019/01/12 23:31:15.886002 [D] dispatching notify.Remove on "/go/example/bar"
2019/01/12 23:31:15.886551 New event /go/example/bar notify.Remove
2019/01/12 23:31:15.886946 New event /go/example/foo notify.Remove
2019/01/12 23:31:16.891939 [D] dispatching notify.Create on "/go/example/foo"
2019/01/12 23:31:16.892981 [D] dispatching notify.Create on "/go/example/bar"
2019/01/12 23:31:16.894072 New event /go/example/foo notify.Create
2019/01/12 23:31:16.894875 New event /go/example/bar notify.Create
2019/01/12 23:31:16.901495 [D] dispatching notify.Remove on "/go/example/foo"
2019/01/12 23:31:16.902762 [D] dispatching notify.Remove on "/go/example/bar"
2019/01/12 23:31:16.903351 New event /go/example/bar notify.Remove
2019/01/12 23:31:16.903970 New event /go/example/foo notify.Remove
2019/01/12 23:31:17.910139 [D] dispatching notify.Create on "/go/example/foo"
2019/01/12 23:31:17.911002 [D] dispatching notify.Create on "/go/example/bar"
2019/01/12 23:31:17.911753 New event /go/example/bar notify.Create
2019/01/12 23:31:17.913244 New event /go/example/foo notify.Create
2019/01/12 23:31:17.921562 [D] dispatching notify.Remove on "/go/example/foo"
2019/01/12 23:31:17.922417 New event /go/example/foo notify.Remove
2019/01/12 23:31:17.923073 [D] dispatching notify.Remove on "/go/example/bar"
2019/01/12 23:31:17.924265 New event /go/example/bar notify.Remove
2019/01/12 23:31:18.906830 [D] dispatching notify.Create on "/go/example/foo"
2019/01/12 23:31:18.906947 [D] dispatching notify.Create on "/go/example/bar"
2019/01/12 23:31:18.907202 New event /go/example/bar notify.Create
2019/01/12 23:31:18.907303 New event /go/example/foo notify.Create
2019/01/12 23:31:18.915407 [D] dispatching notify.Remove on "/go/example/foo"
2019/01/12 23:31:18.915546 New event /go/example/foo notify.Remove
2019/01/12 23:31:18.916539 [D] dispatching notify.Remove on "/go/example/bar"
2019/01/12 23:31:18.916624 New event /go/example/bar notify.Remove
2019/01/12 23:31:19.924494 [D] dispatching notify.Create on "/go/example/foo"
2019/01/12 23:31:19.924959 [D] dispatching notify.Create on "/go/example/bar"
```

When I get the `receiver too slow` message, every event that happens on this folder is not sent
to the channel. I can remove/create `foo` several times, it won't be watched again.