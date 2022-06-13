(set tag (fs.import "constructor.lsp"))
(set tags (of
    (tag "print" (of "\"Hello world!\""))
    (tag "print" (of "\"Hi Jack!\""))
    (tag "print" (of (tag "add" (of "2" "2"))))
))
(print (list.new tags))