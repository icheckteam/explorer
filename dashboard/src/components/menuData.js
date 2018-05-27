// This file is shared across the demos.

import React from 'react';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import { Link } from "react-router-dom";
import { ROUTES } from '../common/constants'
export const menuListItems = (
  <div>
    <Link to={ROUTES.HOME}>
      <ListItem button>
        <ListItemText primary="Home" />
      </ListItem>
    </Link>
    <Link to={ROUTES.BLOCKS}>
      <ListItem button>
        <ListItemText primary="Blocks" />
      </ListItem>
    </Link>
    <Link to={ROUTES.TXS}>
      <ListItem button>
        <ListItemText primary="Transactions" />
      </ListItem>
    </Link>
    <Link to={ROUTES.ASSETS}>
      <ListItem button>
        <ListItemText primary="Assets" />
      </ListItem>
    </Link>
    <Link to={ROUTES.ADDRESSES}>
      <ListItem button>
        <ListItemText primary="Addresses" />
      </ListItem>
    </Link>
  </div>
);
