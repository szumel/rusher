### Rusher

The idea is to create a tool that will streamline and simplify the execution of scripts (with security in mind) necessary to deploy the application.

The idea is to define in the file the steps that must be made in the shell.

File format is xml, for now. (schema.xml in repository root)

In order to fiddling around, just execute make in repo root then ./rusher [--help]

Usage:
rusher -s=filename.xml -e=test




## Change Current Working Directory [changeCwd]
### Changing current working directory to new
Params:
* dir -> New current working directory. There are predefined values you can use: {projectPath}

-------------------------------

## Composer install [composerInstall]
### Install composer dependencies
-------------------------------

## Copy Files [copyFiles]
### Copy files and folders from place to place
Params:
* from -> which file/dir should be copied?
* to -> file/dir name to which copy

-------------------------------

## Git clone [gitClone]
### Cloning git repository
Params:
* origin -> repository origin path
* dir -> to which location clone
* key -> path to ssh key

-------------------------------

## Magento 2 setup upgrade [magento2SetupUpgrade]
### Executes setup:upgrade command on magento 2 intance
-------------------------------

## Magento2 enable modules [magento2EnableModules]
### Enable all modules in magento 2 instance
-------------------------------

## Magento 2 compile [magento2Compile]
### Compiles magento 2 instance
-------------------------------

## Move [move]
### Move file/dir to new location with new name
Params:
* source -> source file/dir path to move
* dst -> destination file/dir path

-------------------------------

## Print PWD [printPwd]
### Print to stdout process working directory
-------------------------------

## Print String [printString]
### Printing to stdout provided string. 
 String should be provided as param 'text'.
Params:
* text -> string

-------------------------------

## Remove Dir [removeDir]
### Removes given directory
Params:
* dir -> directory path which will be removed

-------------------------------

## Symlink [symlink]
### create symlink in context of current working direcory
Params:
* source -> link source
* target -> link target

-------------------------------

## Open Link [openLink]
### open link to warm up cache of website
Params:
* url -> link to open

-------------------------------

## aptInstall [aptInstall]
### install package via apt get
Params:
* package -> package to install
* accept -> accept all dependencies to install [y/n]

-------------------------------
