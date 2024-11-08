insert into veterinarian (
                          full_name,
                          email,
                          phone,
                          password_hash,
                          position,
                          clinic_number
) VALUES('Ivanov Ivan Ivanovich',
         'ivanov@mail.ru',
         '+79998762302',
         'hash_test',
         'veterinarian',
         '892847245451');


insert into owner (full_name, email, phone, password_hash) VALUES ('Vasiliy Ivanovich Chyrkov',
                                                                   'vasilyich@example.com',
                                                                   '+78889087678',
                                                                   'hash_test');


insert into device (unique_number, status) VALUES (
                                                    '12345678', 'WORKING'
                                                      )