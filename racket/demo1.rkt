#lang racket

(define scale 2.5)

(define (inc n)
  (+ 1 n))

(define (good-password x)
  (and (string? x)                   ;; must be a string
       (<= 8 (length x))             ;; at least 8 chars
       (not (string-contains? x " ") ;; has no spaces
            )))

;;
;; Why is this an incorrect implementation for "and"?
;;
(define (bad-and e1 e2)
  (if e1
      (if e2 #t #f)
      #f
      ))

(define (launch-missiles)
  (println "missiles have been launched") ;; print a message
  #t                                      ;; #t signals missiles fired
  )

(define code 'halt)
