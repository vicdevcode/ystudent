database:

- user:
- - id
- - fio
- - email
- - phone
- - membership?

user - user

off_chat_group - creator (admin)
off_chat_group - members (teacher - curator, students)
off_chat_group - admin (teacher - curator, student - leader)

off_chat_subject - creator (teacher)
off_chat_subject - members (teacher, students)
off_chat_subject - admins (teacher, can be any members)

unoff_chat - creator (anyone)
unoff_chat - members (invited users)
unoff_chat - admins (creator, can be any members)
