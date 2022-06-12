(set sss (list (of print)))
(set print (def t (do
    (set log (list-get sss 0))
    (log "-----------------")
    (log t)
    (log "-----------------")
)))

(iterate (range 1 100) i (do
    (print "Hi jack!")
))