-- Enums
CREATE TYPE question_type AS ENUM ('single_choice', 'multiple_choice', 'text', 'matching');
CREATE TYPE attempt_status AS ENUM ('in_progress', 'finished', 'expired');

-- Roles
CREATE TABLE roles (
                       id serial PRIMARY KEY,
                       name varchar(50) UNIQUE NOT NULL
);

-- Departments
CREATE TABLE departments (
                             id serial PRIMARY KEY,
                             name varchar(255) UNIQUE NOT NULL
);

-- Units
CREATE TABLE units (
                       id serial PRIMARY KEY,
                       name varchar(255) UNIQUE NOT NULL
);

-- Users
CREATE TABLE users (
                       id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                       role_id int NOT NULL REFERENCES roles(id),
                       curator_id uuid REFERENCES users(id),
                       unit_id int REFERENCES units(id),
                       department_id int REFERENCES departments(id),
                       full_name varchar(255),
                       practice_start date,
                       practice_end date
);

-- Question Pools
CREATE TABLE question_pools (
                                id serial PRIMARY KEY,
                                name varchar(255) NOT NULL,
                                description text,
                                time_limit_seconds int NOT NULL,
                                created_by uuid NOT NULL REFERENCES users(id),
                                owner_id uuid REFERENCES users(id),
                                unit_id int NOT NULL REFERENCES units(id),
                                created_at timestamp DEFAULT now()
);

-- Questions
CREATE TABLE questions (
                           id serial PRIMARY KEY,
                           pool_id int REFERENCES question_pools(id),
                           text text NOT NULL,
                           question_type question_type NOT NULL DEFAULT 'single_choice',
                           score int NOT NULL DEFAULT 1,
                           position int NOT NULL,
                           media_url text
);

-- Options
CREATE TABLE options (
                         id serial PRIMARY KEY,
                         question_id int NOT NULL REFERENCES questions(id),
                         text text NOT NULL,
                         is_correct boolean NOT NULL DEFAULT false
);

-- Tests
CREATE TABLE tests (
                       id serial PRIMARY KEY,
                       title varchar(255) NOT NULL,
                       description text,
                       created_by uuid NOT NULL REFERENCES users(id),
                       created_at timestamp DEFAULT now()
);

-- Test Pools
CREATE TABLE test_pools (
                            test_id int NOT NULL REFERENCES tests(id),
                            pool_id int NOT NULL REFERENCES question_pools(id),
                            PRIMARY KEY (test_id, pool_id)
);

-- User Groups
CREATE TABLE user_groups (
                             id serial PRIMARY KEY,
                             name varchar(255) NOT NULL
);

-- User Group Members
CREATE TABLE user_group_members (
                                    group_id int NOT NULL REFERENCES user_groups(id),
                                    user_id uuid NOT NULL REFERENCES users(id),
                                    PRIMARY KEY (group_id, user_id)
);

-- Test Group Assignments
CREATE TABLE test_group_assignments (
                                        id serial PRIMARY KEY,
                                        test_id int NOT NULL REFERENCES tests(id),
                                        group_id int NOT NULL REFERENCES user_groups(id),
                                        assigned_by uuid NOT NULL REFERENCES users(id),
                                        assigned_at timestamp DEFAULT now(),
                                        deadline timestamp
);

-- Test Attempts
CREATE TABLE test_attempts (
                               id serial PRIMARY KEY,
                               test_id int NOT NULL REFERENCES tests(id),
                               assignment_id int NOT NULL REFERENCES test_group_assignments(id),
                               user_id uuid NOT NULL REFERENCES users(id),
                               started_at timestamp DEFAULT now(),
                               finished_at timestamp,
                               score int,
                               max_score int,
                               status attempt_status NOT NULL DEFAULT 'in_progress'
);

-- Test Attempt Answers
CREATE TABLE test_attempt_answers (
                                      id serial PRIMARY KEY,
                                      attempt_id int NOT NULL REFERENCES test_attempts(id),
                                      question_id int NOT NULL REFERENCES questions(id),
                                      option_id int REFERENCES options(id),
                                      text text,
                                      is_correct boolean,
                                      answered_at timestamp DEFAULT now()
);
