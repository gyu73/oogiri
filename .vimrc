""""""""""""""""""""""""""""""
" 各種オプションの設定
" """"""""""""""""""""""""""""""
" 行番号を表示する
set number
" 改行時に自動的にインデントが振られるようにする
set autoindent
" tabを半角スペースで挿入する
set expandtab
" Vimが自動で生成するインデントの幅をスペース4つ分にする
set shiftwidth=4
" tabキーを押した時の幅をスペース２つ文にする
set tabstop=2

""""""""""""""""""""""""""""""
" プラグインのセットアップ
" """"""""""""""""""""""""""""""
set nocompatible                " Be iMproved
if has('vim_starting')

  " Required:
  set runtimepath+=~/.vim/bundle/neobundle.vim
endif

  " Required:
  call neobundle#begin(expand('~/.vim/bundle/'))

  " Let NeoBundle manage NeoBundle
  " " Required:
  NeoBundleFetch 'Shougo/neobundle.vim'
  NeoBundle 'scrooloose/syntastic'
call neobundle#end()

" Required:
filetype plugin indent on
