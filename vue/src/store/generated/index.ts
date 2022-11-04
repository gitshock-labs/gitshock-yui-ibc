// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import GitshockClaims from './gitshock.claims'
import GitshockDex from './gitshock.dex'
import GitshockGitshock from './gitshock.gitshock'
import GitshockIncentives from './gitshock.incentives'


export default { 
  GitshockClaims: load(GitshockClaims, 'gitshock.claims'),
  GitshockDex: load(GitshockDex, 'gitshock.dex'),
  GitshockGitshock: load(GitshockGitshock, 'gitshock.gitshock'),
  GitshockIncentives: load(GitshockIncentives, 'gitshock.incentives'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}