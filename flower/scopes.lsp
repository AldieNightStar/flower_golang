(set test (def a b c (do
    (print a b c) ; This should print 10 20 30
    (print etc) ; This should print nil
)))

(set caller (def (do
    (set x 10)
    (set y 20)
    (set z 30)
    (set etc 1000) ; This var should not be in test func scope
    (test x y z)
)))

(caller)