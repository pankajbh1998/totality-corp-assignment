# totality-corp-assignment

Steps to run the application
  1. docker-compose --build
  2. docker-compose up
  
Port = 8001

The End Points are:
  1. :8080/user/1             // for one user
  2. :8080/users/[1,2,3]      // for more than one users 


The application consists of 3 layred architecture.

Tests are written for three layers with 90.6% test coverage

The Mock Database consists of UserIds in range 1 to 5 inclusive.
