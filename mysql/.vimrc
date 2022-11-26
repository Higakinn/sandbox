call plug#begin()
Plug 'mattn/vim-sqlfmt'
" フォーマットのオプションを変更
let g:sqlfmt_program = "sqlformat --comma_first true -r -k upper -o %s -"

" マッピング設定
nmap <buffer><leader>sf <Plug>(sqlfmt)
call plug#end()
