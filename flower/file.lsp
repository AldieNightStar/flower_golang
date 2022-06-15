; Example that shows how to iterate numbers faster
(set it (iteration (range 1 1000)))

(set runner (do
    (iterate (infinite) i (do
        (set n (next it))
        (if (is-nil n) (do (break)))
        (if (eq (mod i 100) 0) (do
            (print (str.concat "Thread-" num ":" (next it)))
        ))
    ))
))

(thread (do (call runner (with "num" 1))))
(thread (do (call runner (with "num" 2))))
(thread (do (call runner (with "num" 3))))
(thread (do (call runner (with "num" 4))))

(time.sleep 10000)