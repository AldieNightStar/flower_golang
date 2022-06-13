(set parse (fs.import "parser.lsp"))
(print (parse "hello 1 2 DDD"))

(iterate (fs.list "./") f (do
    (print f.name)
))