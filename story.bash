#!/bin/sh

printf "\n\nSignup user with nickname Gopher\n"
curl --request PUT --data '{"nickname": "Gopher"}' http://localhost:8083/signup
printf "\n\nGopher asks new question\n"
curl --request PUT --data '{"author_id": 1, "title": "My first question", "content": "How are you?"}' http://localhost:8083/ask
printf "\n\nGopher asks new question\n"
curl --request PUT --data '{"author_id": 1, "title": "My second question", "content": "How old are you?"}' http://localhost:8083/ask
printf "\n\nSignup user with nickname Vasya\n"
curl --request PUT --data '{"nickname": "Vasya"}' http://localhost:8083/signup
printf "\n\nVasya answers first gophers question (he spammed x8)\n"
curl --request PUT --data '{"author_id": 2,"question_id": 1,"content": "Im fine1"}' http://localhost:8083/answer
printf "\n"
curl --request PUT --data '{"author_id": 2,"question_id": 1,"content": "Im fine2"}' http://localhost:8083/answer
printf "\n"
curl --request PUT --data '{"author_id": 2,"question_id": 1,"content": "Im fine3"}' http://localhost:8083/answer
printf "\n"
curl --request PUT --data '{"author_id": 2,"question_id": 1,"content": "Im fine4"}' http://localhost:8083/answer
printf "\n"
curl --request PUT --data '{"author_id": 2,"question_id": 1,"content": "Im fine5"}' http://localhost:8083/answer
printf "\n"
curl --request PUT --data '{"author_id": 2,"question_id": 1,"content": "Im fine6"}' http://localhost:8083/answer
printf "\n"
curl --request PUT --data '{"author_id": 2,"question_id": 1,"content": "Im fine7"}' http://localhost:8083/answer
printf "\n"
curl --request PUT --data '{"author_id": 2,"question_id": 1,"content": "Im fine8"}' http://localhost:8083/answer
printf "\n\nGopher thinks third answer was pretty good. He choosed it as best\n"
curl --request PATCH --data '{"id":3}' http://localhost:8083/answer/best
printf "\n\nVasya answers second gophers question (he spammed again x3)\n"
curl --request PUT --data '{"author_id": 2,"question_id": 2,"content": "Im 11"}' http://localhost:8083/answer
printf "\n"
curl --request PUT --data '{"author_id": 2,"question_id": 2,"content": "Im 12"}' http://localhost:8083/answer
printf "\n"
curl --request PUT --data '{"author_id": 2,"question_id": 2,"content": "Im 13"}' http://localhost:8083/answer
printf "\n\nGopher becames mad and deletes secon question\n"
curl --request DELETE --data '{"id": 2}' http://localhost:8083/question/delete
printf "\n\nThat's all, falks!\n"