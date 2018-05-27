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

function Assets(props) {
  const { classes, assets } = props;
  return (
    <Paper className={classes.root}>
      <Table className={classes.table}>
        <TableHead>
          <TableRow>
            <TableCell>ID</TableCell>
            <TableCell>Name</TableCell>
            <TableCell>Addresses</TableCell>
            <TableCell>Transactions</TableCell>
            <TableCell>Regitered</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {assets.map(n => {
            return (
              <TableRow key={n.id}>
                <TableCell>{n.id}</TableCell>
                <TableCell>{n.name}</TableCell>
                <TableCell>{n.numaddrs}</TableCell>
                <TableCell>{n.numtxs}</TableCell>
                <TableCell>{moment(n.created).fromNow()}</TableCell>
              </TableRow>
            );
          })}
        </TableBody>
      </Table>
    </Paper>
  );
}

Assets.propTypes = {
  classes: PropTypes.object.isRequired,
  assets: PropTypes.array
};

export default withStyles(styles)(Assets);
