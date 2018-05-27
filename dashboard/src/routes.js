import React from 'react'
import { Route, BrowserRouter as Router } from "react-router-dom";
import MenuDrawer from './components/MenuDrawer'

import { ROUTES  } from './common/constants'

// all containers 
import HomeContainer from './containers/HomeContainer'
import BlocksContainer from './containers/BlocksContainer'
import AssetsContainer from './containers/AssetsContainer'
import AddressesContainer from './containers/AddressesContainer'
import TransactionsContainer from './containers/TransactionsContainer'


export const Routes = (props) => {
  return(
    <Router>
      <MenuDrawer>
        <Route path={ROUTES.HOME} exact component={HomeContainer} />
        <Route path={ROUTES.TXS} exact component={TransactionsContainer} />
        <Route path={ROUTES.ADDRESSES} exact component={AddressesContainer} />
        <Route path={ROUTES.ASSETS} exact component={AssetsContainer} />
        <Route path={ROUTES.BLOCKS} exact component={BlocksContainer} />
      </MenuDrawer>
    </Router>
  )
}

export default Routes