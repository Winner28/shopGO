'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _classnames = require('classnames');

var _classnames2 = _interopRequireDefault(_classnames);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

var propTypes = {
  className: _propTypes2.default.string,
  children: _propTypes2.default.node
};

var CardText = function CardText(_ref) {
  var className = _ref.className,
      children = _ref.children;
  return _react2.default.createElement(
    'section',
    {
      className: (0, _classnames2.default)('mdc-card__supporting-text', className)
    },
    children
  );
};
CardText.propTypes = propTypes;
exports.default = CardText;