### Rusher

The idea is to create a tool that will streamline and simplify the execution of scripts necessary to deploy the application.

The idea is to define in the file the steps that must be made in the shell.

File format is xml, for now. (schema.xml in repository root)

In order to fiddling around, just execute make in repo root then ./rusher [--help]

Usage:
rusher -s=filename.xml -e=test
#COMMAND
	Name: Change Current Working Directory
	Code: changeCwd
	Description: Changing current working directory to new
	Params:
		dir -> New current working directory. There are predefined values you can use: {projectPath}

#END

#COMMAND
	Name: Composer install
	Code: composerInstall
	Description: Install composer dependencies

#END

#COMMAND
	Name: Copy Files
	Code: copyFiles
	Description: Copy files and folders from place to place
	Params:
		to -> file/dir name to which copy
		from -> which file/dir should be copied?

#END

#COMMAND
	Name: Git clone
	Code: gitClone
	Description: Cloning git repository
	Params:
		origin -> repository origin path
		dir -> to which location clone
		key -> path to ssh key

#END

#COMMAND
	Name: Magento 2 setup upgrade
	Code: magento2SetupUpgrade
	Description: Executes setup:upgrade command on magento 2 intance

#END

#COMMAND
	Name: Magento2 enable modules
	Code: magento2EnableModules
	Description: Enable all modules in magento 2 instance

#END

#COMMAND
	Name: Magento 2 compile
	Code: magento2Compile
	Description: Compiles magento 2 instance

#END

#COMMAND
	Name: Move
	Code: move
	Description: Move file/dir to new location with new name
	Params:
		source -> source file/dir path to move
		dst -> destination file/dir path

#END

#COMMAND
	Name: Print PWD
	Code: printPwd
	Description: Print to stdout process working directory

#END

#COMMAND
	Name: Print String
	Code: printString
	Description: Printing to stdout provided string. 
 String should be provided as param 'text'.
	Params:
		text -> string

#END

#COMMAND
	Name: Remove Dir
	Code: removeDir
	Description: Removes given directory
	Params:
		dir -> directory path which will be removed

#END

#COMMAND
	Name: Symlink
	Code: symlink
	Description: create symlink in context of current working direcory
	Params:
		source -> link source
		target -> link target

#END

#COMMAND
	Name: Open Link
	Code: openLink
	Description: open link to warm up cache of website
	Params:
		url -> link to open

#END
