set number
set autoindent
set expandtab
set shiftwidth=4
set tabstop=2

set nocompatible
if has('vim_starting')
  set runtimepath+=~/.vim/bundle/neobundle.vim
endif
  call neobundle#begin(expand('~/.vim/bundle/'))
  NeoBundleFetch 'Shougo/neobundle.vim'
  NeoBundle 'scrooloose/syntastic'
call neobundle#end()
filetype plugin indent on

	
