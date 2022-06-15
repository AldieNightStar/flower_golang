(set gen (fs.import "generator.lsp"))
(set parser (fs.import "parser.lsp"))
(set bench (fs.import "benchmark.lsp"))

(iterate gen.multiplied elem (do
    (print elem)
))

(print (parser.parse "Hello world and all"))

(print "BENCHMARK")
(bench)

(print (str.concat "Type of bench is: " (type bench)))