cmake_minimum_required(VERSION 3.1.0 FATAL_ERROR)

set(CMAKE_MODULE_PATH ${CMAKE_CURRENT_LIST_DIR}/)

project(gosk Go)

# peg file
add_custom_target(gosk_peg_go peg src/gosk.peg)

# gosk
set(SOURCE_FILES
    src/gosk.peg.go)
    src/gosk.go)

add_executable(gosk ${SOURCE_FILES})
add_dependencies(gosk gosk_peg_go)
