# ITCS-4102-Project
This repository is for ITCS 4102 Final Project.
Project overview:
       Our project is the implementation of a game of tick-tack-toe in the Go programming langue. This program consists of several classes each with their function contributing to the functioning of the game. 

Class Descriptions:
GUI class - renders the game space the use of buttons and allows the players to take their turns by clicking the appropriate tile.

Main class - keeps track of the state of the game by the use of a 2d array and an integer telling who's turn it is. It also calls the other classes to make alterations to the board state and to check.

Tile class - makes changes to the board by which button was pressed and whose turn it is.

Checks class - checks the board state after each move made and determines if someone has won the game or if it is a draw.

Setup:
first - install the latest version of JVM.
Second - install the golang tool and set up your Gopaths http://www.wadewegner.com/2014/12/easy-go-programming-setup-for-windows/
Third - set up the GUI support. https://github.com/therecipe/qt#installation use the Regular setup section
