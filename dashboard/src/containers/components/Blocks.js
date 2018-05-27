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

function Blocks(props) {
  const { classes, blocks } = props;
  return (
    <Paper className={classes.root}>
      <Table className={classes.table}>
        <TableHead>
          <TableRow>
            <TableCell>Index</TableCell>
            <TableCell>Time</TableCell>
            <TableCell>Transactions</TableCell>
            <TableCell>Validator</TableCell>
            <TableCell>Size</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {blocks.map(n => {
            return (
              <TableRow key={n.height}>
                <TableCell>{n.height}</TableCell>
                <TableCell>{moment(n.time).fromNow()}</TableCell>
                <TableCell>{n.numtxs}</TableCell>
                <TableCell>{n.validator}</TableCell>
                <TableCell>{n.size}</TableCell>
              </TableRow>
            );
          })}
        </TableBody>
      </Table>
    </Paper>
  );
}

Blocks.propTypes = {
  classes: PropTypes.object.isRequired,
  blocks: PropTypes.array
};

export default withStyles(styles)(Blocks);
