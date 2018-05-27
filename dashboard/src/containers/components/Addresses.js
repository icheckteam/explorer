import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import moment from 'moment';
const styles = theme => ({
  root: {
    width: '100%',
    marginTop: theme.spacing.unit * 3,
    overflowX: 'auto',
  },
  table: {
    minWidth: 700,
  },
});

function Addresses(props) {
  const { classes, addresses } = props;
  return (
    <Paper className={classes.root}>
      <Table className={classes.table}>
        <TableHead>
          <TableRow>
            <TableCell>Address</TableCell>
            <TableCell>Created</TableCell>
            <TableCell>Last Transaction</TableCell>
            <TableCell>Transactions</TableCell>
            <TableCell>Assets</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {addresses.map(n => {
            return (
              <TableRow key={n.address}>
                <TableCell>{n.address}</TableCell>
                <TableCell>{moment(n.created).fromNow()}</TableCell>
                <TableCell>{moment(n.last_tx).fromNow()}</TableCell>
                <TableCell>{n.numtxs}</TableCell>
                <TableCell>10ICN</TableCell>
              </TableRow>
            );
          })}
        </TableBody>
      </Table>
    </Paper>
  );
}

Addresses.propTypes = {
  classes: PropTypes.object.isRequired,
  addresses: PropTypes.array
};

export default withStyles(styles)(Addresses);
