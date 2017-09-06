""""""""""""""""""""""""""""""
" 各種オプションの設定
" """"""""""""""""""""""""""""""
" 行番号を表示する
set number
" 改行時に前の行のインデントを継続する
set autoindent
" タブ入力を複数の空白入力に置き換える
set expandtab
" Vimが挿入するインデントの幅指定
set shiftwidth=4
" タブ文字の表示幅を指定
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
