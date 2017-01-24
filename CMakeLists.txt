cmake_minimum_required(VERSION 3.1.0 FATAL_ERROR)

set(CMAKE_MODULE_PATH ${CMAKE_CURRENT_LIST_DIR}/)

project(gosk Go)

# peg file
set(gosk_peg ${gosk_SOURCE_DIR}/src/gosk.peg)

add_custom_target(gosk_peg_go
  COMMAND peg ${gosk_peg}
)

add_custom_target(gosk_yacc
  COMMAND go tool yacc -o ${gosk_SOURCE_DIR}/src/parser.go ${gosk_SOURCE_DIR}/src/parser.go.y
)

# gosk
set(SOURCE_FILES
  ${gosk_SOURCE_DIR}/src/parser.go
  ${gosk_SOURCE_DIR}/src/gosk.peg.go
  ${gosk_SOURCE_DIR}/src/gosk.go
)

if(MSYS)
  set(gosk_BIN ${gosk_SOURCE_DIR}/gosk.exe)
else()
  set(gosk_BIN ${gosk_SOURCE_DIR}/gosk)
endif()

add_custom_target(gosk
  COMMAND go build -o ${gosk_BIN} ${SOURCE_FILES}
)

add_dependencies(gosk gosk_peg_go)
add_dependencies(gosk gosk_yacc)
