(set parse (fs.import "parser.lsp"))
(print (parse "hello 1 2 DDD"))

(iterate (fs.list "./") f (do
    (dict.set f "rura" (dict.new (with "x" 1)))
    (print f)
))