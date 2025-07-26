#!/bin/bash

gentool -dsn "root:123456@tcp(127.0.0.1:3306)/blog" \
  -tables "comments,posts,users" \
  -outPath models/gen \
  -fieldSignable

