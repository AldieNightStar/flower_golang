(set gen (fs.import "generator.lsp"))
(set parser (fs.import "parser.lsp"))

(iterate gen.multiplied elem (do
    (print elem)
))

(print (parser.parse "Hello world and all"))