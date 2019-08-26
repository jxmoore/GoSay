package models

import (
	"math/rand"
	"time"
)

// getGopher is a function that returns a raw string litteral of a random gopher represented in ASCII.
func getGopher() string {

	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	switch random.Intn(7) {
	case 0:
		return `
                                          '.--://++++++++///::-.'                                    
                               '.:/+syhdmmmmmmmmmmmmmmmmmmmmmmmmmdhyo+:.                            
                          ':+shddmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmddyo/.                       
                      ':ohdmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmdddddmmmmmmdho-      .--:-.'       
       '.::::-.    '/ydmmmmdhyysssyydmmmmmmmmmmmmmmmmmmmmmmhsssssyyyyssssshmmmmdh+..shdddddddy/.    
    '/ydddmmmddh--sdmmmdsssshdmmmmmdysssymmmmmmmmmmmmmmmhooymNMMMmdmNMMMNmyosmmmmmds/ymmmmmmmmmdo'  
  '+dmmmmmmmmms/ydmmmyosdNMMMMMdyyhNMMMNhosdmmmmmmmmmmdosmMMMMMh-'''.oNMMMMNh/hmmmmmdo/++oymmmmmmy' 
 'ymmmmdsooyh/sdmmmd+yNMMMMMMm-'   .oMMMMMd/hmmmmmmmmh/dMMMMMMd'      oMMMMMMN/ymmmmmmh-  'ommmmmms 
 smmmmy.    :hmmmmd/dMMMMMMMM:       dMMMMMN/hmmmmmmh:NMMMMMMMd   /-  +MMMMMMMN:dmmmmmmd/  :mmmmmmm 
 dmmmm/    :dmmmmm/dMMMMMMMMM/  -o' 'mMMMMMMm:mmmmmm/dMMMMMMMMMs.'y+'/NMMMMMMMMsommmmmmmd+-dmmmmmmh 
 ymmmmd/. :dmmmmmd:MMMMMMMMMMNo-/y-:dMMMMMMMM:dmmmmm.MMMMMMMMMMMNdhdmMMMMMMMMMMd/mmmmmmmmd:ymmmmmd- 
 .hmmmmms-dmmmmmmh/MMMMMMMMMMMMNmmmMMMMMMMMMM/hmmmmm.MMMMMMMMMMMMMMMMMMMMMMMMMMh/mmmmmmmmmh-dmmds.  
  '+hmmd-hmmmmmmmd-MMMMMMMMMMMMMMMMMMMMMMMMMM-dmmmmmosMMMMMMMMMMMMMMMMMMMMMMMMM/ymmmmmmmmmmo+s/.    
    '-+/+mmmmmmmmm+yMMMMMMMMMMMMMMMMMMMMMMMMsommmmmmd/hMMMMMMMMMMMMMMMMMMMMMMNoommmmmmmmmmmd.       
       'dmmmmmmmmmd/hMMMMMMMMMMMMMMMMMMMMMNs+mmmmdddddosmMMMMMMMMMMMMMMMMMMMd+smmmmmmmmmmmmmo       
       +mmmmmmmmmmmdoomMMMMMMMMMMMMMMMMMNh+ymdy+:-.--:ososdNMMMMMMMMMMMMNmhosdmmmmmmmmmmmmmmd'      
       hmmmmmmmmmmmmmhsohmNMMMMMMMMMNmdsoydmh-         'ydyssyhddmmddhyssshmmmmmmmmmmmmmmmmmm/      
      'mmmmmmmmmmmmmmmmdyysssyyyyysssyydmdhs:          ':yhmmdhyyyyyyhdmmmmmmmmmmmmmmmmmmmmmmy      
      :mmmmmmmmmmmmmmmmmmmmmmdddddmmmmmmy/oyh+-.'''..:+yhyo+ydmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmd      
      /mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmh-hdddddddhhdddddddddo/dmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm-     
      +mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmo/dddddddddddddddddddd:ymmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm/     
      ommmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmd:shhhhyoo+-sooshhddho/dmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmo     
      ommmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmyo:oydNMh-MMMm+:oooymmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmms     
      +mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm-mMMMMs/MMMMh+mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmy     
      +mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm+yMMMMs/MMMMd/mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmh     
      :mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmd:dMMd/:dMMMoommmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmd     
      -mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmdsoosmdsoooymmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmd     
`

	case 1:
		return `                                            
                                             \ʕ◔ϖ◔ʔ/`

	case 2:
		return `
  
                                                  ,_---~~~~~----._         
                                            _,,_,*^____      _____''*g*\"*, 
                                          / __/ /'     ^.   /     \ ^@q   f 
                                          [  @f | @))   |   | @)) !   0 _/  
                                          \'/   \~____ / __ \_____/    |   
                                            |           _!__|_         !   
                                            }          [______]        !  
                                            ]            | | |         |  
                                            ]             ~ ~          |  
                                            |                          |   
                                            |                          |   
  `

	case 3:
		return `
  
                                             '..---:::---..''                                         
                                    '.:/osyyyyyyhhhhhyyyyyyso/-.                                    
                           '.'   .:oyyyssyyhhhhhhhhhhhyyyyyyysyys/..:/+/:'                          
                        .+syyyo:oysyydmNmmhyyyhhhhhhsydNNNMNNNdsyhyoosshhy/                         
                       :yho/+osyysmNMMMMMMMMNhshhhhomMMMMMMMMMMNyyhhy/'/hhh:                        
                       yhy' -yhhoNdssdMMMMMMMMsyhhomd/-/mMMMMMMMMohhhy/+hhh/                        
                       /yhs-yhhhsN' .-mMMMMMMMdohh+M+ '+oMMMMMMMMohhhhy/hy+'                        
                        -++shhhhoN+-:oNMMMMMMMyyhhyymo/yNMMMMMMMdohhhhhs-'                          
                          -hhhhhysmNNMMMMMMMNysso+++smNMMMMMMNmyshhhhhhh-                           
                          ohhhhhhysydmmmmmdyys/'    -syyhhhhyyyyhhhhhhhho                           
                          yhhhhhhhhhyyyyyyyyohy+::/+yhsyhhhhhhhhhhhhhhhhy                           
                          hhhhhhhhhhhhhhhhhssmmdddddmmd+hhhhhhhhhhhhhhhhh.                          
                          hhhhhhhhhhhhhhhhhhssydm/mdoysshhhhhhhhhhhhhhhhh-                          
                          hhhhhhhhhhhhhhhhhhhshMM+MMohhhhhhhhhhhhhhhhhhhh:                          
                          yhhhhhhhhhhhhhhhhhhhohyoyhohhhhhhhhhhhhhhhhhhhh/                          
                          ohhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh/                          
                          +hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh/                          
                          :hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh/                          
                          -hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh/                          
                          .hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh+                          
                      '-/+:hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh+oo/.                      
                     :hmmm+hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhoymmy:                     
                      oyo-.hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhy'/o+                      
                          -hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh                          
                          /hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh'                         
                          +hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh.                         
                          ohhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh.                         
                          shhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh.                         
                          yhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh.                         
                          yhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhy                          
                          shhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhho                          
                          +hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh:                          
                          -hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhs'                          
                           ohhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhy-                           
                           'shhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhy:                            
                            .ohhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhy:                             
                             '/shhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhyo.                              
                             ':sssyhhhhhhhhhhhhhhhhhhhhhhhhhhhhhyysyh+'                             
                            /ymmmhs/+syyyhhhhhhhhhhhhhhhhhhyys+:.:hmmhh.                            
                            yomdy/'   '.-:/+oosssssssoo+/:-.'     'odh+-                            
                             :/-                                    '.                              
  
  `

	case 4:
		return `
                                                                               
                                                '--::-::--.       '--.                              
                                          .:://+//oo+/+ooo++/:-./++//+s+-                           
                                      '-://::::+o/.'     '-+o::sy+soo/::/s                          
                                   '-/+::::::/h-         '--:y::/ymdNd:::y'                         
                                '-:/:::::::::y.         +mNNysy:::/ody:::s                          
                             ':oo/::::::::::/o          sNMmy+y::::::+o+o.                          
                           ./oo++++++//::::::y          ':++-'y::::::::/s-                          
                         -/+oo:.'   ''.:+/:::o:              :o::::::::::o+                         
                       '++/o:'       -++/oo:::++.          .++::::::::::::++                        
                      .s/:s'        /mMNyys/:::///::-...-:///::::::::::::::+o                       
                     .h/::y         :mNNyo.s:/+sss+++////:::::::::::::::::::s+                      
                 '://+h:::y          .--. .yymNmNNhooss/:::::::::::::::::::::y-                     
                +s/::/y:::o-             'ooNNNmds+++++h:::::::::::::::::::::/s'                    
               :m:::oho::::o/.         ':oossoo+++++ooss::::::::::::::::::::::++                    
               +y::+mNo:::::/+/:....--/++:yo++++o+s--s+::::::::::::::::::::::::+o'                  
               .d/::/ss:::::::://////:::::/soosy' --  +/::::::::::::::::::::::::+y:                 
                :yo///s::::::::::::::::::::://:o-  -/.-o:::::::::::::::::::::::::/d/                
                 '.-:/h:::::::::::::::::::::::::o:-:s+/:::::::::::::::::::::::::::/d.'''''          
                      -+:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::os-...:/         
                       +/:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::y-'  /+         
                        s/::::::::::::::::::::::::::::::::::::::::::::::::::::::::::+m:--:'         
                        'o:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::N:             
                         's::::::::::::::::::::::::::::::::::::::::::::::::::::::::::dy             
                          .y:::::::::::::::::::::::::::::::::::::+//////:::::::::::::om             
                           'o/:::::::::::::::::::::::::::::::::::s     -o:::::::::::::N'            
                             /o::::::::::::::::::::::::::::::::/os::/' ss:::::::::::::d/            
                              :h/:::::::::::::::::::::::::::::/o:::::++/::::::::::::::y+            
                               -d+::::::::::::::::::::::::::::::::::::::::::::::::::::s/            
                                'y/:::::::::::::::::::::::::::::::::::::::::::::::::::s-            
                                 'y:::::::::::::::::::::::::::::::::::::::::::::::::::s:            
                                  s+::::::::::::::::::::::::::::::::::::::::::::::::::s+            
                                  :m::::::::::::::::::::::::::::::::::::::::::::::::::s/            
                                   do:::::::::::::::::::::::::::::::::::::::::::::::::h-            
                                   /d:::::::::::::::::::::::::::::::::::::::::::::::::d             
                                   .d::::::::::::::::::::::::::::::::::::::::::::::::/d             
                                    h::::::::::::::::::::::::::::::::::::::::::::::::s+             
                                    y::::::::::::::::::::::::::::::::::::::::::::::::d.             
                                    y:::::::::::::::::::::::::::::::::::::::::::::::+y              
                                    d:::::::::::::::::::::::::::::::::::::::::::::::y:..'           
                                   'd::::::::::::::::::::::::::::::::::::::::::::::s: '.::          
                                    y::::::::::::::::::::::::::::::::::::::::::::/ys..  /y          
                                    o::::::::::::::::::::::::::::::::::::::::::/oo-'./:./-          
                                    .y:::::::::::::::::::::::::::::::::::::::/++.     ''            
                                    +:o:::::::::::::::::::::::::::::::::::/++:'                     
                                    --:o/::::::::::::::::::::::::::::://o+:.                        
                                        :+/:::::::::::::::::::::://+o+/-                            
                                          :oo/+/++/:::::::::/+ooo+:.                                
                                            's'  '/y++oo+//:/-'                                     
                                           '+'   -+.                                                
                                           s.:  -+                                                  
                                           ::s/+o                                                   
                                                                                                    
  `
	case 5:
		return `
                          
                                                                        .//+/'                        
                                          -+++/'                       oo--:/y.                       
                                        -y---/y/.:///++++oo++++ooooo+y+...-:oo                       
                                        .y--/+ss/:----.............-:d.s.-::d-                       
                                          +ho/--.....................:h-mos+yss+-                     
                                        .o+:..........................+yy+ss/::/oo-                   
                                      /s:.............................-:/:-:::::/s+'                 
                                      +o-...................................-::::::os'                
                                    -y............-//+++++++/:-.............-::::::+y'               
                                    h-..........:++:.'    '-:+s+-............-:::://+s               
                                '--:/h/:.......-o+.          '.-+y/-...........::::o+:y/              
                              -+/:.'''-y:......s/              ..:h/-..........-::::+//h'             
                            '++'       /y...../s               '..:h:..........-:::::::s/             
                            o/         -h.....y-                ...d:-.........-:::::::/h             
                          'h          .d.....s:                ...d:-..........::::::::d'            
                          -y          'd.....:y'     '-/:.     ..+s:-..........::::::::y-            
                            h'      -/s/y/::::-+o'   .+/MMN:   '.+y:-...........::::::::s+            
                            -s.    .y+NMhNMMMMs:/s:' /mmMMMo  ':so:-............::::::::oo            
                            .o+.'  +yy+oNNNmyosss/o+/ohdho.:+so:-.............-::::::/+yy            
                              ./+++///sm///:::/+d-..-/++++/:-.................-:::/sso/.d            
                                    '''+yssysssyh/............................-:oso+-....h'           
                                :+ooo++m:.d'/'+y..........................:oo+/.........y-           
                                .d:::::::h-+oy+y:.....................-+ohd+'     .......s:           
                                :ss/::::+y......................:+ooo/+s/'       '......o:           
                      .//:         '/ossosNo..............-:++++od+o+:s+.          ......s:           
                    +hsyhs           '-hohd+......-:/osydo:.'''y:''-++-.          ......y-           
                    /hssyym'          :s' ':syooshdhsdysyyd+//+/:.-d-..-:++/      '......d'           
                  .dssyyhms/:/+::++:/o+'   .sh+:+o+yhssyyd+' ':+o+y/    .-d      '.....-h            
                  yyssydhssdyssmysydm:o:::/dsshdsshdssyyyhd/ys+::::y:''./s:      ......++            
                  -dssyyddyymhyymhyyds.....-dyyhmyymsssyyyyd d/::/+osho+o:'      .......h'            
                  hsssyyhdddyyddsyddy'    '/sddyydmhssyyyyyN -/+od/:-..'        '....../s             
                -dssyyhdooyhoohyooym/:/+/syoodsoomssyyyyyyN'  '/o             '......-h'             
                yyssyyydyshhsshhsshdyysssyhssdyshhssyyyyyhm+:+os'      -.    '.......s:              
                .dssyyyhdddNyddyyddmhs---:shddhhdNssyyyyyyymo::-'       :/'   .......+o               
                oyssyyymsssmssydssymmms++yhssdyshdssyyyyyyymo++oo/:.'    '   '......+o                
                dssyyyyddhhNhyydhyhhsoy+-/dhyddhmsssyyyyyyym/---:/++o+:.     ......+o'                
              -dssyyyhyssmysydssyd////sysoshhoymssyyyyyyyhm/..-::::::+oo:  '...--o+                  
              syssyyym/:/d::+d::+dyooooyy::ys:yhssyyyyyyyym:.:so+oss/-:/yo----::+s''                 
              dsssyyyhhyymyossssso-....-ossyssmssyyyyyyyyyN::h/:::/sd+oo+//oso+/so....'              
              'mssyyyhhyyyd:...............-..:mssyyyyyyyyyN+d/:::+yo:-.......-:::.''''               
            ''+ooo+:--///:...................ohssyyyyyyyyyhoo:::os:....'''''                         
                              '''''''........shsshhhyhhyyd+y+/os/''                                  
                                        '''''-+o++:.''.--.'./:.                                      
                                                                                                      
    `
	case 6:
		return `
                                      '..-.'                                                        
                              '.-:/+oossssyy:                                                       
                         '-:/osooooooooosyysy-                                                      
                     '-/oooooooooooooosyysssss                                                      
                  .:+oooooooooooooossssssssssh'                                                     
               ./ooooooooooooooossssossssssssy:                                                     
             -+ooooooooooooossssooooossssssssss                                                     
           '+oooooooooosssssooooooooosssssssssh.                                                    
          '+ooossssssyysoooooooooooooossssssssyo                                                    
           :yssssssssssoooooooooooooossssssysssh.                                                   
            .+ysssssssooooooooooooooossssssssssyo                                                   
              -ossssssooooooooooooooossssssssssyd.                                                  
               ':sssssooooooooooooooosssssssssssys'                                                 
                 '/ssssooooooooooooooosssssssssys/o:/++//:'                                         
                  ':ssooooooooooooooooossssssso/:-/hyyyyyss'                                        
                    -ssoooooooooooooooosssso+:-----yysssysy-                                        
                     -ssooooooooooooooso+/--.----/yysssssss'                                        
                      .ssooooooosso+/:---....-:+syssssssso.                                         
                       .ossooo+/:-.........-/ossssssoosy+//::.'                                     
                        'o+::-........--/+ossssssyssss+:------::.'                                  
                       '/ss:------::/+syysssssyyyyso/---........---..-..'''                         
                      -sssssoo+oosssssssssyysssoo/::::::-.........-+:.'''''.'                       
                     'ssyssssssssssssyyyysso/-''       '.::....''..-/.      :y/                     
                     '+syyyssssyyysssso+//o:'-/+/.        :/-..'''..-/      'hy                     
                      '.:s///////::----.-:/'-mddNh'        -/-....-:+s+:     /'                     
                        /-........''''...+. .yddd+         '+-...-syyyyy.   '/           '.         
                       '+.''''''''''''''.+.  '...          'o-..:-.:::-./-'.-'            -'.       
                       '/.'''''''''''''''-/'               +/-.::'  ''''/+-'              '''       
                        /.'''''''''''''''':/'            '/+-...+::::+s/-      :...''  --           
                        :.'''''''''''''''''.:-.'      '.:/:-.....-/--oh.      :---...--''           
                        .:'''''''''''''''''''...---:::::-...'''....-://-     -.''''.../.            
                       ./o.'''''''''''''''''''''''.......'''''''''....--    .-'''''''':             
                       +////---..'''''''''''''''''''''''''''''........:-'' ':------:::.             
                       .o+//////+/:--..'''''''''''''''''''''.-:::--:/+ooo+ -:s:::::+:.              
                       '++/+/::::://////--..'''''''''''''''./++syyhsydyyos -/:...-::.               
                       :o::::-::::::::::///+-..''''''''''''./++syhdhddhys: -/:..:::'                
                     '/o:-----------::::::://+:.'''''''''''':sosysooosos.  ':/-::-                  
                    'o+:--------------::::::::/+.'''''''''''.-:::-:-/o/s.' '/-:.'                   
                    ++----------------::::/+/::/o:-----------......-s/+/..:+:+                      
                   's:--------------::o/://y/::/o/''              '.s/o--:o:+'                      
                   'o:--------:------:/s/::s::::/s.'              '//+'  ::+.                       
                    +-----------------://.-+:::::s-'              -o+.''-:::                        
                    /------------------:://::::::s.'             -o+/:-::'o.                        
                    //-------------------::::::::o'          ''':s-  ....''-:'                      
                    'o----------------------:::://'       '...../:      ''..:.                      
                     /:::-:--------------------:o'         '''.:-                                   
                     -o::-::----:::----------:/o.''''''.--:::+'                                     
                     -y::::::::-:-------:::::os/:::::://:--:+'                                      
                     .y:::::::::::-:-:::::/++:--.....-----//'                                       
                     /+:::/+++//:://///////--........---:+.                                         
                    .y///-''':/---::-............--::+//:/'                                         
                   :o///'     /.'.:/---------:::::-. '/:-::                                         
                  -//:.       /-.-/                    '..'''''                                     
                              '---''''                                                         
  `
	default:
		return "uh oh"
	}

}
