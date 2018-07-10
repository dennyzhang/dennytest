# Suggested Golang package: Logrus
https://github.com/sirupsen/logrus

```
go get "github.com/sirupsen/logrus"
```

| Summary      | Code Snippets|
|:------------|-------------|
| Logging with different level | [log_basic.go](log_basic.go) |
| Logging to file | [log_file.go](log_file.go). Then `ls -lth /tmp/output.txt` |
| Logging to file with log rotation enforced | Not implemented yet: [log_file_rotate.go](log_file_rotate.go) Then `ls -lth /tmp/output*.txt` |


# Discusssions For possible Candiates

- Logrus (logging): https://github.com/sirupsen/logrus

```
the most imported repos for logging

GitHub star: 7,846 stars
From link: https://medium.com/@_orcaman/most-imported-golang-packages-some-insights-fb12915a07
```

- op/go-logging: https://github.com/op/go-logging

```
GitHub star: 1,269 stars

Feedback:
- Doesn't support logging to file
- Need to configure many stuff, before using it
```

- golang/glog: https://github.com/golang/glog

```
GitHub star: 1,852 stars

Feedbacks:
- Too primitive
```
