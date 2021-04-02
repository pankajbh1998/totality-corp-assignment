# totality-corp-assignment

Steps to run the application
  1. docker-compose --build
  2. docker-compose up
  
Port Address = 8001

The End Points are:
  1. :8001/user/1             // for single user
  2. :8001/users/[1,2,3]      // for multiple users 


The application consists of 3 layred architecture.

Tests are written for three layers with 90.6% test coverage

The Mock Database consists of UserIds in range 1 to 5 inclusive.
