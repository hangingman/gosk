cmake_minimum_required(VERSION 3.1.0 FATAL_ERROR)

set(CMAKE_MODULE_PATH ${CMAKE_CURRENT_LIST_DIR}/)

project(gosk Go)

# peg file
set(gosk_peg ${gosk_SOURCE_DIR}/src/gosk.peg)

add_custom_target(gosk_peg_go
  COMMAND peg ${gosk_peg}
)

# gosk
set(SOURCE_FILES
  ${gosk_SOURCE_DIR}/src/gosk.peg.go
  ${gosk_SOURCE_DIR}/src/gosk.go
)

add_custom_target(gosk
  COMMAND go build -o ${gosk_SOURCE_DIR}/gosk ${SOURCE_FILES}
)
add_dependencies(gosk gosk_peg_go)
